/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:29 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RgAuthCtx struct {
	AuthResult AuthResult `json:"authResult"`
	Supi       string     `json:"supi,omitempty"`
	AuthInd    *bool      `json:"authInd,omitempty"`
}
