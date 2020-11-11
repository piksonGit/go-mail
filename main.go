package main

import (
	"fmt"
	"mail/send"
)

func main() {
	var to string
	var subject string
	var content string
	fmt.Println("请输入接收者邮箱：")
	fmt.Scanln(&to)
	fmt.Println("请输入邮件主题：")
	fmt.Scanln(&subject)
	fmt.Println("请输入邮件内容：")
	fmt.Scanln(&content)
	send.SendMail(to, subject, content)
}
