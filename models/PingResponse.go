/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Nov 27 11:12:35 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingResponse struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Nonce   int64  `json:"nonce"`
	From    string `json:"from"`
}
