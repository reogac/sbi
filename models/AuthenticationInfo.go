/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
}
