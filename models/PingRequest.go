/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:23 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingRequest struct {
	Nonce   int64  `json:"nonce"`
	Time    string `json:"time"`
	Message string `json:"message"`
}
