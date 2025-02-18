/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 11 17:13:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingRequest struct {
	Message string `json:"message"`
	Nonce   int64  `json:"nonce"`
	Time    string `json:"time"`
}
