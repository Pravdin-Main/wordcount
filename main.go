package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var src string
	reader := bufio.NewReader(os.Stdin)
	src, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	if err := scanner.Err(); err != nil {
	// 		fmt.Sprintln(fmt.Errorf("Some trouble here, somthing like %s", err))
	// 		return
	// 	}
	// } else {
	// 	fmt.Sprintln("Nothing to scann(((")
	// 	return
	// }
	// var count int
	// for var frase string
	// for _, err := fmt.Scan(&src); err != nil; {
	// 	if err == nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		count++
	// 		_, err = fmt.Scan(&src)
	// 	}
	// }
	// fmt.Fscan(os.Stdin, src)
	// fmt.Scan(&src)
	// fmt.Println(fmt.Scan(&src))
	fmt.Println(len(strings.Fields(src)))
	// fmt.Println(src)
	// fmt.Println(len(strings.Fields(scanner.Text())))
}
