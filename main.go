package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var src string
	// var src_cache string

	// -------- os lib ------------
	// src := os.Args

	// ---------- bufio Scanner lib -----------
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

	// ---------- bufio Reader lib ---------
	// reader := bufio.NewReader(os.Stdin)
	// src, err := reader.ReadString('\n')
	// if err != nil {
	// 	// log.Fatal(err)
	// }
	// fmt.Scan(&src)
	// flag.StringVar(&src, "p", "", "pattern to match against")
	// // flag.Parse()
	// src := strings.Join(flag.Args(), "")

	// ----------- fmt lib ------------
	// for _, err := fmt.Scanln(&src); err != nil; {
	// 	_, err = fmt.Scanln(&src_cache)
	// 	src = src + " " + src_cache
	// 	src_cache = ""
	// }
	// fmt.Fscan(os.Stdin, src)
	// fmt.Scan(&src)
	// fmt.Println(fmt.Scan(&src))
	// fmt.Println(len(strings.Fields(src)))

	// ----------- flag lib --------------
	flag.Parse()
	src = strings.Join(flag.Args(), "")

	// fmt.Println(src)
	fmt.Println(len(strings.Fields(src)))
}
