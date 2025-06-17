/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConsumerNfInformation struct {
	NfId    string `json:"nfId,omitempty"`
	NfSetId string `json:"nfSetId,omitempty"`
	TaiList []Tai  `json:"taiList,omitempty"`
}
