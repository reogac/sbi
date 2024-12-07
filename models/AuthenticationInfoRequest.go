/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:26 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoRequest struct {
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	AusfInstanceId        string                 `json:"ausfInstanceId"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
}
