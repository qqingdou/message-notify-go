package main

import (
	"github.com/qqingdou/message-notify-go/messagenotify"
)

func main()  {

	//you project id and company key
	messagenotify.NewMessageNotify(1, "3d19a744dcf06a48fb591f73af20dfc6")

	//Auto Catch Exception
	defer messagenotify.AutoCatchException()

	yy	:=	0
	t	:=	1 / yy

	println(t)

	messageBody	:=	messagenotify.MessageBody{}
	messageBody.SetTile("test")
	messageBody.SetMessage("test2")
	messageBody.SetType(1)

	messagenotify.GetInstance().AddMessage(messageBody).Push()

	messageBody2	:=	messagenotify.MessageBody{}
	messageBody2.SetTile("test")
	messageBody2.SetMessage("test3")
	messageBody2.SetType(1)
	messagenotify.GetInstance().AddMessage(messageBody2).Push()
}
