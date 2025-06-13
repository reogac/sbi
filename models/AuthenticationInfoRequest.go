/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoRequest struct {
	AusfInstanceId        string                 `json:"ausfInstanceId"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
}
