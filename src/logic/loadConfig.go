package logic


import (
    "io/ioutil"
    "encoding/json"
	"mailsvr/model"
	"mailsvr/data"
	"mailsvr/def"
)


/**
*   parse json config
*/
func ParseSendMailConfig(fileName string) {
    resultConfig := model.SendMailConfig{}
    rawData, err := ioutil.ReadFile(fileName)
    if nil != err {
        return
    }
    dataJson := []byte(rawData)
    err = json.Unmarshal(dataJson, &resultConfig)
    if nil != err {
        return
    }
    data.GlobalConfig = resultConfig
}


/**
*    load send mail config
*/
func LoadConfig() int {
	ParseSendMailConfig("conf/config.json")

    // listen port < 0, then throw error
    if data.GlobalConfig.ListenPort <= 0 {
        return def.ERROR_INTERNAL
    }
    return def.RETURN_NORMAL
}