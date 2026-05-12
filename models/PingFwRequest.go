/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:23 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingFwRequest struct {
	Time    string `json:"time"`
	Service string `json:"service"`
	Message string `json:"message"`
	Nonce   int64  `json:"nonce"`
}
