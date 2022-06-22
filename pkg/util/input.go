package util

import (
	"bufio"
	"fmt"
	"os"
)

var stdin *bufio.Scanner

func ReadLine() string {
	if stdin == nil {
		stdin = bufio.NewScanner(os.Stdin)
	}

	stdin.Scan()
	if err := stdin.Err(); err != nil {
		panic(err)
	}
	return stdin.Text()
}

func Prompt(s string) string {
	fmt.Println(s)
	return ReadLine()
}
