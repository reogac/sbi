/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PwsInformation struct {
	OmcId             string            `json:"omcId,omitempty"`
	NfId              string            `json:"nfId,omitempty"`
	MessageIdentifier int               `json:"messageIdentifier"`
	SerialNumber      int               `json:"serialNumber"`
	PwsContainer      N2InfoContent     `json:"pwsContainer"`
	BcEmptyAreaList   []GlobalRanNodeId `json:"bcEmptyAreaList,omitempty"`
	SendRanResponse   *bool             `json:"sendRanResponse,omitempty"`
}
