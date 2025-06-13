/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionManagementSubscriptionData struct {
	SingleNssai                     Snssai                             `json:"singleNssai"`
	DnnConfigurations               map[string]DnnConfiguration        `json:"dnnConfigurations,omitempty"`
	SharedVnGroupDataIds            map[string]string                  `json:"sharedVnGroupDataIds,omitempty"`
	SharedDnnConfigurationsId       string                             `json:"sharedDnnConfigurationsId,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	SharedTraceDataId               string                             `json:"sharedTraceDataId,omitempty"`
	SuggestedPacketNumDlList        map[string]SuggestedPacketNumDl    `json:"suggestedPacketNumDlList,omitempty"`
	InternalGroupIds                []string                           `json:"internalGroupIds,omitempty"`
	OdbPacketServices               OdbPacketServices                  `json:"odbPacketServices,omitempty"`
	ExpectedUeBehavioursList        map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`
	ThreeGppChargingCharacteristics string                             `json:"3gppChargingCharacteristics,omitempty"`
	SupportedFeatures               string                             `json:"supportedFeatures,omitempty"`
}
