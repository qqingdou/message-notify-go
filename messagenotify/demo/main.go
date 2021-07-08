package main

import (
	"github.com/qqingdou/message-notify-go/messagenotify"
)

func main()  {

	//you project id and company key
	messageNotifyInstance	:=	messagenotify.NewMessageNotify(1, "xxxx")

	messageBody	:=	messagenotify.MessageBody{}
	messageBody.SetTile("first error")
	messageBody.SetMessage("first error")
	messageBody.SetType(1)

	messageNotifyInstance.AddMessage(messageBody).Push()

	messageBody2	:=	messagenotify.MessageBody{}
	messageBody2.SetTile("second error")
	messageBody2.SetMessage("second error")
	messageBody2.SetType(1)
	messagenotify.GetInstance().AddMessage(messageBody2).Push()

	//Auto Catch Exception
	defer messagenotify.AutoCatchException()

	panic("Hello World, I'm a error.")
}
