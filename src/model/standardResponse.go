package model


/**
 *  standard response
 */
type StandardResponse struct {
	Ret int `json:"ret"`
	Info []string `json:"info,omitempty"`
	Msg string `json:"msg,omitempty"`
}