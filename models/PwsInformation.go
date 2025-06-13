/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PwsInformation struct {
	SendRanResponse   *bool             `json:"sendRanResponse,omitempty"`
	OmcId             string            `json:"omcId,omitempty"`
	NfId              string            `json:"nfId,omitempty"`
	MessageIdentifier int               `json:"messageIdentifier"`
	SerialNumber      int               `json:"serialNumber"`
	PwsContainer      N2InfoContent     `json:"pwsContainer"`
	BcEmptyAreaList   []GlobalRanNodeId `json:"bcEmptyAreaList,omitempty"`
}
