package messagenotify

import (
	"crypto/hmac"
	"math/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
	"crypto/md5"
	"encoding/hex"
	"crypto/aes"
	"bytes"
	"encoding/base64"
	"crypto/sha256"
)

type MessageNotify struct {
	projectId	int
	key			string
	notifyUrl	string
	messages	[]map[string]string
}

var myMessageNotify *MessageNotify

func GetInstance() *MessageNotify {
	return myMessageNotify
}

func NewMessageNotify(projectId int, key string) *MessageNotify {
	myMessageNotify	=	&MessageNotify{projectId: projectId, key: key}
	return	myMessageNotify
}

func (messageNotify *MessageNotify) AddMessage(messageBody MessageBody) *MessageNotify {
	messageNotify.messages = append(messageNotify.messages, messageBody.toMap())
	return	messageNotify
}

func (messageNotify *MessageNotify)Push() string {
	messages		:=	messageNotify.messages
	sendData, err 	:=	json.Marshal(messages)
	if err != nil {
		return ""
	}

	messageNotify.clear()

	return messageNotify.notify(string(sendData))
}

func (messageNotify *MessageNotify)GetNotifyUrl() string {
	notifyUrl	:=	messageNotify.notifyUrl
	if len(notifyUrl) > 0 {
		return notifyUrl
	}else{
		return "https://open-api.51baocuo.com/message/notify"
	}
}

func (messageNotify *MessageNotify)SetNotifyUrl(notifyUrl string) {
	messageNotify.notifyUrl	=	notifyUrl
}

func (messageNotify *MessageNotify)GetProjectId() int  {
	return messageNotify.projectId
}

func (messageNotify *MessageNotify)SetProjectId(projectId int)  {
	messageNotify.projectId	=	projectId
}

func (messageNotify *MessageNotify)SetKey(key string)  {
	messageNotify.key	=	key
}

func (messageNotify *MessageNotify)getKey() string {
	return messageNotify.key
}

func (messageNotify *MessageNotify)notify(message string) string {
	nowUnixTime	:=	time.Now().Unix()
	nonce		:=	getUniqueId(fmt.Sprintf("%d", messageNotify.projectId))
	header		:=	"Content-Type: application/json"

	aesEncryptResult,err	:= aesEncrypt(message, messageNotify.getAesKey())

	if err != nil {
		return	""
	}

	params					:=	make(map[string]string)
	params["time"]			=	fmt.Sprintf("%d", nowUnixTime)
	params["nonce"]			=	nonce
	params["project_id"]	=	fmt.Sprintf("%d", messageNotify.GetProjectId())
	params["messages"]		=	aesEncryptResult
	signStr					:=	buildSignStr(params)
	params["sign"]			=	hashHmac(signStr, messageNotify.getKey())
	jsonBuffer, err			:=	json.Marshal(params)

	if err != nil {
		return ""
	}

	sendResult	:=	post(messageNotify.GetNotifyUrl(), string(jsonBuffer), 5, header)

	return sendResult
}

func post(url string, params string, timeout int, contentType string) string {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Post(url, contentType, bytes.NewBuffer([]byte(params)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func (messageNotify *MessageNotify)getAesKey() string {
	key		:=	messageNotify.getKey()
	keyLength	:=	16
	if len(key) > keyLength {
		return	key[0:keyLength]
	}else{
		return fmt.Sprintf("%16s#0", key)
	}
}

func (messageNotify *MessageNotify)clear()  {
	messageNotify.messages	=	make([]map[string]string, 0)
}

func buildSignStr(params map[string]string) string {
	var keys	[]string
	for key := range params {
		keys	=	append(keys, key)
	}
	sort.Strings(keys)
	var sortedString	[]string
	for _, key := range keys{
		sortedString	=	append(sortedString, fmt.Sprintf("%s=%s", key, params[key]))
	}
	return	strings.Join(sortedString, "&")
}

func getUniqueId(pref string) string {
	nanoSecond	:=	time.Now().UnixNano()
	nonce1		:=	fmt.Sprintf("%08v", rand.New(rand.NewSource(nanoSecond)).Int63n(100000000))
	nonce2		:=	fmt.Sprintf("%08v", rand.New(rand.NewSource(nanoSecond)).Int63n(100000000))
	md5Str		:=	myMd5(fmt.Sprintf("%s%s%s%s", pref, nonce1, nanoSecond, nonce2))
	return md5Str
}

func aesEncrypt(data string, myKey string) (string, error) {
	key			:= []byte(myKey)
	origData	:=	[]byte(data)
	cipher, err	:= aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	length 		:= (len(origData) + aes.BlockSize) / aes.BlockSize
	plain 		:= make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func myMd5(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func hashHmac(data string, key string) string {
	hash:= hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}