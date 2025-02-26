package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetUserInputRequired(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	input, _ := reader.ReadString('\n')
	if len(input) <= 0 {
		GetUserInputRequired(message)
	}
	return strings.TrimSpace(input)
}
