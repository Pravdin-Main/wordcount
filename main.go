package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Sprintln(fmt.Errorf("Some trouble here, somthing like %s", err))
			return
		}
	} else {
		fmt.Sprintln("Nothing to scann(((")
		return
	}
	fmt.Println(len(strings.Fields(scanner.Text())))
}
