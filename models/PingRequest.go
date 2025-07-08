/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:24 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingRequest struct {
	Time    string `json:"time"`
	Message string `json:"message"`
	Nonce   int64  `json:"nonce"`
}
