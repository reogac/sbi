/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PwsInformation struct {
	MessageIdentifier int               `json:"messageIdentifier"`
	SerialNumber      int               `json:"serialNumber"`
	PwsContainer      N2InfoContent     `json:"pwsContainer"`
	BcEmptyAreaList   []GlobalRanNodeId `json:"bcEmptyAreaList,omitempty"`
	SendRanResponse   *bool             `json:"sendRanResponse,omitempty"`
	OmcId             string            `json:"omcId,omitempty"`
	NfId              string            `json:"nfId,omitempty"`
}
