package sensor

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestArduinoReader(t *testing.T) {

	const testString1 = "Hello"
	const testString2 = "World"
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	reader := bytes.NewBufferString(fmt.Sprintf("%s\r\n%s\r\n", testString1, testString2))
	messages := make(chan string)
	go receiveArduinoMessages(ctx, messages, ioutil.NopCloser(reader))

	msg := <-messages
	if testString1 != msg {
		t.Fatalf("expected %v, but received: %v", testString1, msg)
	}
	msg = <-messages
	if testString2 != msg {
		t.Fatalf("expected %v, but received: %v", testString2, msg)
	}
}

func TestArduinoReaderIsStoppedByContext(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*200)
	messages := make(chan string)

	receiveArduinoMessages(ctx, messages, slowReader{})
}

type slowReader struct{}

func (slowReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (slowReader) Close() error {
	return nil
}
