package model


/**
 *  receiver struct
 */
type ReceiverInfo struct {
	Type string `json:"type"`
	List []string `json:"list"`
}


/**
*   config struct
*/
type SendMailConfig struct {
    ListenPort int `json:"listen_port"`
    MailServer string `json:"mail_server"`
    MailPort int `json:"mail_port"`
    SenderAccount string `json:"sender_account"`
    SenderPassword string `json:"sender_password"`
	Environment string `json:"environment"`
	LogServerIp string `json:"log_server_ip"`
	LogServerPort int `json:"log_server_port"`
    ReceiverInfos []*ReceiverInfo `json:"receiver_infos"`
}