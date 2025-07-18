/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingFwRequest struct {
	Message string `json:"message"`
	Nonce   int64  `json:"nonce"`
	Time    string `json:"time"`
	Service string `json:"service"`
}
