package messagenotify

import (
	"fmt"
	"runtime"
)

var AutoCatchException = func() {
	if r := recover(); r != nil {
		messageBody	:=	MessageBody{}
		messageBody.SetTile(fmt.Sprintf("%s", r))
		messageBody.SetMessage(fmt.Sprintf("%s", stack()))
		messageBody.SetType(1)
		GetInstance().AddMessage(messageBody).Push()
	}
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}