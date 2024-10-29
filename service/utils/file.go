package utils

import (
	"github.com/gin-gonic/gin"
	"image"
	_ "image/gif"  // Import image formats to support GIF, JPEG, and PNG
	_ "image/jpeg" // We use underscore imports to register the image formats
	_ "image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func HandleFile(c *gin.Context) (string, string) {
	path, err := filepath.Abs("./")

	file, _ := c.FormFile("file")
	filename := filepath.Base(file.Filename)
	var basePath = filepath.Join(path, "assets")

	if _, err = os.Stat(basePath); os.IsNotExist(err) {
		var dirMod uint64
		if dirMod, err = strconv.ParseUint("0775", 8, 32); err == nil {
			err = os.Mkdir(basePath, os.FileMode(dirMod))
		}
	}
	dst := filepath.Join(basePath, filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)
	return dst, filename
}

func GetFileMimeType(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use http.DetectContentType to determine the file's MIME type
	mimeType := http.DetectContentType(buffer)
	return mimeType, nil
}

func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}

	// Use the Size() method to get the file size in bytes
	return fileInfo.Size(), nil
}

func IsImageMimeType(mimeType string) bool {
	return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "image/gif"
}

func GetImageDimensions(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// Decode the image to get the dimensions
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	// Get the width and height of the image
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	return width, height, nil
}