/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Nov 27 11:35:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingFwRequest struct {
	Service string `json:"service"`
	Message string `json:"message"`
	Nonce   int64  `json:"nonce"`
	Time    string `json:"time"`
}
