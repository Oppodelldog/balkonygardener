package sensor

import (
	"context"
	"fmt"
	"io"

	"github.com/Oppodelldog/balkonygardener/log"
)

func receiveArduinoMessages(ctx context.Context, arduinoMessages chan string, reader io.ReadCloser) {
	const readBufferSize = 100

	log.Info("reading arduino data...")

	readBuffer := make([]byte, readBufferSize)
	var messageBuffer []byte
	for {
		select {
		case <-ctx.Done():
			fmt.Println("closing arduino reader")
			err := reader.Close()
			if err != nil {
				log.Errorf("error closing reader: %v", err)
			}
			return
		default:
			n, err := reader.Read(readBuffer)
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Error(err)
				break
			}
			readBytes := readBuffer[:n]
			messageBuffer = append(messageBuffer, readBytes...)
			var eol int
			var index int
			for index < len(messageBuffer) {
				b := messageBuffer[index]
				if b == '\r' {
					eol++
				}
				if b == '\n' {
					eol++

					if eol == 2 {
						info := string(messageBuffer[:index-1])
						arduinoMessages <- info
						messageBuffer = messageBuffer[index+1:]
						index = 0
						eol = 0
					}
				}
				index++
			}
		}
	}
}
