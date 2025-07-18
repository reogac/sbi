/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingResponse struct {
	Nonce   int64  `json:"nonce"`
	Time    string `json:"time"`
	From    string `json:"from,omitempty"`
	Message string `json:"message"`
}
