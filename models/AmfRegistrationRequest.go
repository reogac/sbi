/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:24 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationRequest struct {
	Uuid   string `json:"uuid"`
	AmfSet string `json:"amfSet"`
	PlmnId PlmnId `json:"plmnId"`
}
