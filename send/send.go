package send

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(from string, to string, text string) {
	auth := smtp.PlainAuth(
		"",
		"no-reply@infinitynewtab.com",
		"wsfRWKJJN86G4BtS",
		"smtp.exmail.qq.com",
	)
	fmt.Println("验证构建完成")
	err := smtp.SendMail(
		"smtp.exmail.qq.com:465",
		auth,
		"no-reply@infinitynewtab.com",
		[]string{"403393082@qq.com"},
		[]byte("To: 403393082@qq.com\r\nFrom: no-reply@infinitynewtab.com\r\nSubject: theme\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\nHello World"),
	)
	fmt.Println("发送完成")
	if err != nil {
		fmt.Println("真的是有错误")
		log.Fatal(err)
		fmt.Println(err)
	}
}
