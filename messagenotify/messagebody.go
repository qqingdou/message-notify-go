package messagenotify

import (
	"encoding/json"
	"fmt"
	"time"
)

type MessageBody struct {
	title			string
	_type			int
	line			string
	file			string
	time			int
	message			string
	requestUrl		string
	requestBody		string
	customData		string
	userAgent		string
	clientIp		string
	requestId		string
}

func (messageBody *MessageBody)GetTitle() string  {
	return	messageBody.title
}

func (messageBody *MessageBody)SetTile(title string)  {
	messageBody.title	=	title
}

func (messageBody *MessageBody)GetType() int  {
	return messageBody._type
}

func (messageBody *MessageBody)SetType(_type int)  {
	messageBody._type	=	_type
}

func (messageBody *MessageBody)GetLine() string  {
	return messageBody.line
}

func (messageBody *MessageBody)SetLine(line string)  {
	messageBody.line	=	line
}

func (messageBody *MessageBody)GetFile() string  {
	return messageBody.file
}

func (messageBody *MessageBody)SetFile(file string)  {
	messageBody.file	=	file
}

func (messageBody *MessageBody)GetTime() int  {
	currTime	:=	messageBody.time
	if currTime <= 0 {
		return int(time.Now().Unix())
	}
	return messageBody.time
}

func (messageBody *MessageBody)SetTime(time int)  {
	messageBody.time	=	time
}

func (messageBody *MessageBody)GetMessage() string  {
	return messageBody.message
}

func (messageBody *MessageBody)SetMessage(message string)  {
	messageBody.message	=	message
}

func (messageBody *MessageBody)GetRequestUrl() string {
	return messageBody.requestUrl
}

func (messageBody *MessageBody)SetRequestUrl(requestUrl string)  {
	messageBody.requestUrl	=	requestUrl
}

func (messageBody *MessageBody)GetRequestBody() string {
	return messageBody.requestBody
}

func (messageBody *MessageBody)SetRequestBody(requestBody string)  {
	messageBody.requestBody	=	requestBody
}

func (messageBody *MessageBody)GetCustomData() string  {
	return messageBody.customData
}

func (messageBody *MessageBody)SetCustomData(customData string)  {
	messageBody.customData	=	customData
}

func (messageBody *MessageBody)GetUserAgent() string {
	return messageBody.userAgent
}

func (messageBody *MessageBody)SetUserAgent(userAgent string)  {
	messageBody.userAgent	=	userAgent
}

func (messageBody *MessageBody)GetClientIp() string {
	return messageBody.clientIp
}

func (messageBody *MessageBody)SetClientIp(clientIp string)  {
	messageBody.clientIp	=	clientIp
}

func (messageBody *MessageBody)GetRequestId() string {
	return messageBody.requestId
}

func (messageBody *MessageBody)SetRequestId(requestId string)  {
	messageBody.requestId	=	requestId
}

func (messageBody *MessageBody)toMap() map[string]string  {
	myMap	:=	map[string]string{
		"type"			:	fmt.Sprintf("%d", messageBody.GetType()),
		"title"			:	messageBody.GetTitle(),
		"file"			:	messageBody.GetFile(),
		"line"			:	messageBody.GetLine(),
		"message"		:	messageBody.GetMessage(),
		"request_url"	:	messageBody.GetRequestUrl(),
		"request_body"	:	messageBody.GetRequestBody(),
		"time"			:	fmt.Sprintf("%d", messageBody.GetTime()),
		"user_agent"	:	messageBody.GetUserAgent(),
		"client_ip"		:	messageBody.GetClientIp(),
		"request_id"	:	messageBody.GetRequestId(),
		"custom_data"	:	messageBody.GetCustomData(),
	}
	return myMap
}

func (messageBody *MessageBody)string() string {
	myMapData	:=	messageBody.toMap()
	messageBodyBuffer,err	:=	json.Marshal(myMapData)
	if err != nil {
		return ""
	}
	return string(messageBodyBuffer)
}