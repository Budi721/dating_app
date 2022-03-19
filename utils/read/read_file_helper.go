package read

import (
	"bufio"
	"github.com/Budi721/dating_app/utils/logger"
	"os"
)

func ReadFile(path string) []byte {
	img, err := os.Open(path)
	defer img.Close()

	if err != nil {
		logger.Log.Error().Err(err)
		return nil
	}
	fileInfo, _ := img.Stat()
	var size int64 = fileInfo.Size()
	buf := make([]byte, size)

	// read content into buffer
	fileReader := bufio.NewReader(img)
	fileReader.Read(buf)
	return buf
}
