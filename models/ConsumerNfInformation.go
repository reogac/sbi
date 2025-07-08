/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConsumerNfInformation struct {
	TaiList []Tai  `json:"taiList,omitempty"`
	NfId    string `json:"nfId,omitempty"`
	NfSetId string `json:"nfSetId,omitempty"`
}
