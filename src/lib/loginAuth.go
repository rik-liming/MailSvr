package lib

import(
    "errors"
)

/*
	auth login
 */
type loginAuth struct {
	username, password string
	host                         string
}
 
/*
	auth login 验证
 */
func LoginAuth(username, password, host string) Auth {
	return &loginAuth{username, password, host}
}
 
/*
	初步验证服务器信息，输入账号
 */
func (a *loginAuth) Start(server *ServerInfo) (string, []byte, error) {
	// 如果不是安全连接，也不是本地的服务器，报错，不允许不安全的连接
	if !server.TLS && !isLocalhost(server.Name) {
		return "", nil, errors.New("unencrypted connection")
	}
	// 如果服务器信息和 Auth 对象的服务器信息不一致，报错
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	// 验证时需要的账号
	resp := []byte(a.username)
	// "auth login" 命令
	return "LOGIN", resp, nil
}
 
/*
	进一步进行验证，输入密码
 */
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	// 如果服务器需要更多验证，报错
	if more {
		return []byte(a.password), nil
	}
	return nil, nil
}
