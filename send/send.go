package send

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(to string, subject string, text string) {
	auth := smtp.PlainAuth(
		"",
		"life@qijing.fun",
		"USb7N9uKda5upc8f",
		"smtp.exmail.qq.com",
	)
	fmt.Println("验证构建完成")
	body := fmt.Sprintf("To:%s\r\nFrom:life@qijing.fun\r\nSubject:%s\r\nContent-type:text/html;charset=UTF-8\r\n\r\n%s", to, subject, text)
	err := smtp.SendMail(
		"smtp.exmail.qq.com:587",
		auth,
		"life@qijing.fun",
		[]string{to},
		[]byte(body),
	)
	fmt.Println("发送完成")
	if err != nil {
		fmt.Println("真的是有错误")
		log.Fatal(err)
		fmt.Println(err)
	}
}
