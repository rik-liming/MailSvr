package logic


import(
	"fmt"
	"mailsvr/def"
	"mailsvr/protocol"
	"mailsvr/lib"
	"mailsvr/data"
)

/**
 *  record log content
 */
func RecordLog(game string, content string) {
	// create tcp connection
	conn, err := lib.TcpConnect(data.GlobalConfig.LogServerIp, uint16(data.GlobalConfig.LogServerPort))
	if nil != err {
		fmt.Println(err)
	}
	
	// send tcp log
	pack := protocol.NewWriteObj()
	pack.Init(def.CMD_LOG_SVR_WRITE)
	pack.WriteString("mailsvr")
	pack.WriteString(game)
	pack.WriteString(content)
	lib.Send(conn, pack.GetBuf())
	lib.Close(conn)
}