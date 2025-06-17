/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:57 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
}
