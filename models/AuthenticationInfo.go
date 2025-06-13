/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	TraceData             *TraceData             `json:"traceData,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	Pei                   string                 `json:"pei,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
}
