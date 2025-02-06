/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Feb  6 13:48:14 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationRequest struct {
	Uuid   string `json:"uuid"`
	AmfSet string `json:"amfSet"`
	PlmnId PlmnId `json:"plmnId"`
}
