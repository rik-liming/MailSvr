package logic


import (
	"fmt"
	"strings"
	"strconv"
	"encoding/base64"
	"mailsvr/def"
	"mailsvr/lib"
	"mailsvr/data"
)


/**
*    send mail function
*/
func SendMail(title string, content string, game string) int {
	
	mailServer := data.GlobalConfig.MailServer
	senderAccount := data.GlobalConfig.SenderAccount
	senderPassword := data.GlobalConfig.SenderPassword
	receiverInfos := data.GlobalConfig.ReceiverInfos
	mailPort := data.GlobalConfig.MailPort
	environment := data.GlobalConfig.Environment
	
	// find the receiver list
	receiverList := []string{}
	for _, receiverInfo := range receiverInfos {
		if receiverInfo.Type == game {
			receiverList = receiverInfo.List
			break
		}
	}
	
	// if receiver list is empty, return
	if len(receiverList) == 0 {
		fmt.Println("receiver list is empty!")
		return def.ERROR_ILLEGAL_PARAM
	}

	// assemble mail format
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	header := make(map[string]string)
	header["To"] = strings.Join(receiverList, ",")
	fullTitle := "[env]: " + environment + ", [game]: " + game + ", " + title
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(fullTitle)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(content))
	
	// record log
	logContent := fullTitle + " [content]: " + content
	RecordLog(game, logContent)
	
	// send mail
	mailAuth := lib.LoginAuth(senderAccount, senderPassword, mailServer)
	for _, receiver := range receiverList {
		err := lib.SendMail(mailServer + ":" + strconv.Itoa(mailPort), mailAuth, senderAccount, []string{receiver}, []byte(message))
		if err != nil {
			fmt.Println(err)
		}
	}
	
	return def.RETURN_NORMAL
}