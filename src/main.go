package main


import (
    "fmt"
    "mailsvr/logic"
    "mailsvr/def"
	"mailsvr/controller"
)


/**
*    the main function
*/
func main() {
	// load config
    configCode := logic.LoadConfig()

    // result code not normal, then return
    if configCode != def.RETURN_NORMAL {
	fmt.Println("read config.json error, please check!")
        return
    }
	
	// start http server
    serverCode := controller.StartSendMailServer()
    if serverCode != def.RETURN_NORMAL {
    fmt.Println("start http server error, please check!")
        return
    }
}
