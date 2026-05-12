/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:41 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
}
