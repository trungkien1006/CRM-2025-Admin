package helpers

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

func StoreFile(image *multipart.FileHeader) error {
	// Lưu file vào hệ thống
	file, err := image.Open()
	if err != nil {
		return errors.New("Lỗi mở file")
	}
	defer file.Close()

	dst, err := os.Create("public/images/" + image.Filename)
	if err != nil {
		return errors.New("Lỗi lưu file")
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return errors.New("Lỗi ghi file")
	}

	return nil
}