package controller


import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"mailsvr/def"
	"mailsvr/data"
	"mailsvr/logic"
	"mailsvr/model"
)

/**
 *   start the send mail server
 */
func StartSendMailServer() int {
	http.HandleFunc("/mail/send", sendMailAction)
	err := http.ListenAndServe(":" + strconv.Itoa(data.GlobalConfig.ListenPort), nil)
	if nil != err {
		fmt.Println(err.Error())
		return def.ERROR_INTERNAL
	}
	return def.RETURN_NORMAL
}

/**
*    send mail action
*/
func sendMailAction(w http.ResponseWriter, r *http.Request) {
	
	// data init
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	game := r.PostFormValue("game")
	
	// param error
	if title == "" || content == "" || game == "" {
		result := model.StandardResponse {
			Ret: def.ERROR_ILLEGAL_PARAM,
			Info: []string{},
			Msg: "param error",
		}
		resultJson, err := json.Marshal(result)
		if nil != err {
			fmt.Println(err)
		}
		w.Write(resultJson)
		return
	}
	
	// start send mail
	resultCode := logic.SendMail(title, content, game)
	result := model.StandardResponse {
		Ret: resultCode,
		Info: []string{},
		Msg: "",
	}
	resultJson, err := json.Marshal(result)
	if nil != err {
		fmt.Println(err)
	}
	w.Write(resultJson)
}