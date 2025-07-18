/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoRequest struct {
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	AusfInstanceId        string                 `json:"ausfInstanceId"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
}
