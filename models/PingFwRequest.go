/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Feb 10 19:19:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PingFwRequest struct {
	Nonce   int64  `json:"nonce"`
	Time    string `json:"time"`
	Service string `json:"service"`
	Message string `json:"message"`
}
