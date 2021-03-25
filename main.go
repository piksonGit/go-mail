package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"mail/send"
	"os"
)

func main() {
	var to string
	var subject string
	var content string
	fmt.Println("请输入接收者邮箱：")
	fmt.Scanln(&to)
	fmt.Println("请输入邮件主题：")
	fmt.Scanln(&subject)
	tpl := template.Must(
		template.ParseFiles("./send/body.html"),
	)
	htmlstring := tpl.ExecuteTemplate(os.Stdout, "warning", map[string]interface{}{
		"Time": "2021.03.25 14:00",
	})
	fmt.Println(htmlstring)
	bytes, error := ioutil.ReadFile("./send/body.html")
	var bodyString string
	fmt.Println(bodyString)
	if error != nil {
		fmt.Println(error)
	} else {
		bodyString = string(bytes)
	}
	content = bodyString
	send.SendMail(to, subject, content)
}
