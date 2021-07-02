package main

import (
	"bytes"
	"fmt"
	"html/template"
	"mail/send"
	"os"
)

func main() {
	var subject string
	//io.Write是通过流的形式写。所以不一定一次会全部写入
	//第一个参数标题，第二个参数正文，第三个到最后一个的参数是要发送的邮件们。
	//第一个参数应该为级别，info,warning, error
	b := new(bytes.Buffer)
	args := os.Args
	var tempKey int
	var emails []string
	for k, v := range args {
		if v == "|||" {
			tempKey = k + 1
			break
		}

	}
	//template.HTML可以阻止golang转义html标签
	htmlstring := template.HTML(args[2])
	emails = args[tempKey:]
	//to = args[1]
	subject = args[1]

	tpl := template.Must(
		template.ParseFiles("./body.html"),
	)

	err := tpl.ExecuteTemplate(b, "warning", map[string]interface{}{
		"Detail": htmlstring,
	})
	if err != nil {
		fmt.Println(err)
	}

	send.SendMail(emails, subject, b.String())

}
