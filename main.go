package main

import "mail/send"
import "fmt"

func main() {

	fmt.Println("啥子情况开始了")
	send.SendMail("no-reply@infinitynewtab.com","403393082@qq.com","我是你的小苹果")
	fmt.Println("到底是啥子情况")
}