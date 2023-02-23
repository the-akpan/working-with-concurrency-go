package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(test *testing.T) {
	var wg sync.WaitGroup
	message := "Message"

	wg.Add(1)
	go updateMessage(message, &wg)
	wg.Wait()

	if msg != message {
		test.Error("`msg` was not updated")
	}
}

func Test_printMessage(test *testing.T) {
	stdOut := os.Stdout
	read, write, _ := os.Pipe()

	os.Stdout = write

	printMessage()

	write.Close()
	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, msg) {
		test.Errorf("Output does not contain %s", msg)
	}
}

func Test_main(test *testing.T) {
	stdOut := os.Stdout
	read, write, _ := os.Pipe()

	os.Stdout = write

	main()

	write.Close()
	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	errors := make([]string, 0)
	for _, word := range msgs {
		if !strings.Contains(output, word) {
			errors = append(errors, fmt.Sprintf("%s not found", word))
		}
	}

	if len(errors) != 0 {
		test.Error(strings.Join(errors, "\n"))
	}
}
