/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingFwRequest struct {
	Time    string `json:"time"`
	Service string `json:"service"`
	Message string `json:"message"`
	Nonce   int64  `json:"nonce"`
}
