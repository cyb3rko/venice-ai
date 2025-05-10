package utils

import (
	"bufio"
	"encoding/base64"
	"io"
	"os"
	"strings"
)

func PersistBase64ToFile(id string, format string, data string) error {
	data = "data:image/png;base64," + data
	i := strings.Index(data, ",")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))

	if format == "" {
		format = "webp"
	}
	file, err := os.Create(tmpDataPath + id + "." + format)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}
	return writer.Flush()
}
