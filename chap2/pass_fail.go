package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin) // 從STDIN讀取
	input, err := reader.ReadString('\n') // 看到\n符號就停止讀取
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input) // 去除換行符號或空白
	grade, err := strconv.ParseFloat(input, 64) // 字串轉型
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(status)
}
