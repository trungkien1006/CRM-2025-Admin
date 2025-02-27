package helpers

import "time"

func GetCurrentTimeVN() time.Time {
	location, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	currentTime := time.Now().In(location)

	return currentTime
}

func IsPast(dateStr string) (bool, error) {
	// Định dạng thời gian theo yêu cầu: yyyy-MM-ddThh:mm
	const layout = "2006-01-02T15:04"

	// Chuyển đổi chuỗi thành kiểu time.Time
	inputTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return false, err
	}

	// Lấy thời gian hiện tại
	now := time.Now()

	// So sánh với thời gian hiện tại
	return inputTime.Before(now), nil
}