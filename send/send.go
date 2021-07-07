package send

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
)

func SendMail(to []string, subject string, text string) {
	smptinfo := ReadJson("./smtpinfo.json")
	auth := smtp.PlainAuth(
		"",
		smptinfo["email"],
		smptinfo["password"],
		smptinfo["host"],
	)
	fmt.Println("验证构建完成")
	body := fmt.Sprintf("To:%s\r\nFrom:%s\r\nSubject:%s\r\nContent-type:text/html;charset=UTF-8\r\n\r\n%s", to, smptinfo["email"], subject, text)
	err := smtp.SendMail(
		smptinfo["hostandport"],
		auth,
		smptinfo["email"],
		to,
		[]byte(body),
	)
	fmt.Println("发送完成")
	if err != nil {
		fmt.Println("真的是有错误")
		log.Fatal(err)
		fmt.Println(err)
	}
}

func ReadJson(filepath string) map[string]string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取json配置错误")
	}
	var jsonContent map[string]string = make(map[string]string)
	// Unmarshal第二个参数必须是指针
	err = json.Unmarshal(data, &jsonContent)
	if err != nil {
		fmt.Println("json解析出错", err)
	}
	return jsonContent
}
