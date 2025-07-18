/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfo struct {
	TraceData             *TraceData             `json:"traceData,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	FetchUeAmData         *bool                  `json:"fetchUeAmData,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
}
