/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	SupiOrSuci            string                 `json:"supiOrSuci"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
}
