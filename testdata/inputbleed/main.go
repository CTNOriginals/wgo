package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("\n-- start --\n")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")

	for scanner.Scan() {
		cmd := strings.TrimSpace(scanner.Text())

		fmt.Printf("< %s\n", cmd)
		fmt.Printf("> ")
	}
}
