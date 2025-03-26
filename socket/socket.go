package main

import (
	"admin-v1/app/helpers"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var connStore = make(map[int]*websocket.Conn)
var mu sync.Mutex

type receiverMessage struct {
	ReceiverId 	int
	Message		string
} 

type senderMessage struct {
	SenderId	int
	Message		string
} 

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { 
		fmt.Println("Yêu cầu từ Origin:", r.Header.Get("Origin")) // Debug origin
		return true 
	}, // Chấp nhận mọi origin (có thể điều chỉnh)
}

func handleWebSocket(c *gin.Context) {
	header := c.GetHeader("Sec-WebSocket-Protocol")

	fmt.Println("Header nhận được:", header)

	if err := helpers.CheckJWT(header); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	idUser := helpers.GetTokenSubject(header).Id

	fmt.Println("Nhận yêu cầu nâng cấp WebSocket...")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, http.Header{
    "Sec-WebSocket-Protocol": []string{header}, // Phản hồi lại subprotocol
})

	if err != nil {
		fmt.Println("Lỗi khi nâng cấp WebSocket:", err)
		return
	}

	fmt.Println("Kết nối WebSocket thành công!")

	mu.Lock()
	connStore[int(idUser)] = conn
	mu.Unlock()

	go ReadMessageHandler(conn, int(idUser))
}

func ReadMessageHandler(conn *websocket.Conn, senderId int) {
	fmt.Println("Bắt đầu lắng nghe tin nhắn từ client:", senderId)  // Log kiểm tra

	defer func() {
		fmt.Printf("⚠️ Đóng kết nối WebSocket cho user %d\n", senderId)
		mu.Lock()
		delete(connStore, senderId)
		mu.Unlock()
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Đóng kết nối bình thường"))
		conn.Close()
	}()

	for {
		fmt.Println("Đang chờ nhận message...")
		
		// Đọc message từ client
		_, messageReceiveJson, err := conn.ReadMessage()

		if err != nil {
			mu.Lock()
			if connStore[senderId] != nil {
				delete(connStore, senderId)
			}
			mu.Unlock()

			fmt.Println("Client ngắt kết nối: " + err.Error())
			break
		}

		var messageReceive receiverMessage

		err = json.Unmarshal([]byte(messageReceiveJson), &messageReceive)

		if err != nil {
			fmt.Println("Lỗi khi parse json:", err)

			break
		}
		
		fmt.Println("Nhận:", string(messageReceiveJson))

		// mu.Lock()
		// if err := connStore[senderId].WriteMessage(websocket.TextMessage, []byte("Bạn vừa gửi 1 tin nhắn!")); err != nil {
		// 	fmt.Println("Không thể gửi message, có thể kết nối đã đóng:", err)
		// 	mu.Unlock()
		// 	break
		// }
		// mu.Unlock()

		mu.Lock()
		if connStore[messageReceive.ReceiverId] != nil {
			var senderMessage = senderMessage{
				SenderId: senderId,
				Message: messageReceive.Message,
			}

			fmt.Println("Bắt đầu gửi tin nhắn cho receiver")

			senderMessageJson, err := json.Marshal(senderMessage)
			if err != nil {
				fmt.Println("Lỗi khi mã hóa thành json:", err)
				mu.Unlock()
				break
			}

			if err := connStore[messageReceive.ReceiverId].WriteMessage(websocket.TextMessage, []byte(senderMessageJson)); err != nil {
				fmt.Println("Không thể gửi message, có thể kết nối đã đóng:", err)
				mu.Unlock()
				break
			}

			// if err := connStore[senderId].WriteMessage(websocket.TextMessage, []byte("Gửi tin nhắn thành công cho người nhận")); err != nil {
			// 	fmt.Println("Không thể gửi message, có thể kết nối đã đóng:", err)
			// 	mu.Unlock()
			// 	break
			// }
		}
		mu.Unlock()
		
	}
}

func main() {
	if os.Getenv("DOCKER_ENV") != "true" {
		_ = godotenv.Load("../.env") // Chỉ tải .env nếu không chạy trong Docker
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	// Route xử lý WebSocket
	r.GET("/ws", handleWebSocket)

	port := os.Getenv("PORT_SOCKET")

	ln, err := net.Listen("tcp", "0.0.0.0:" + port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, r)
}