/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoRequest struct {
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	AusfInstanceId        string                 `json:"ausfInstanceId"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
}
