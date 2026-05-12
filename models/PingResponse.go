/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:23 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingResponse struct {
	Message string `json:"message"`
	Nonce   int64  `json:"nonce"`
	Time    string `json:"time"`
	From    string `json:"from,omitempty"`
}
