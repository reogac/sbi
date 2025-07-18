/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionManagementSubscriptionData struct {
	SupportedFeatures               string                             `json:"supportedFeatures,omitempty"`
	OdbPacketServices               OdbPacketServices                  `json:"odbPacketServices,omitempty"`
	SharedTraceDataId               string                             `json:"sharedTraceDataId,omitempty"`
	InternalGroupIds                []string                           `json:"internalGroupIds,omitempty"`
	SharedVnGroupDataIds            map[string]string                  `json:"sharedVnGroupDataIds,omitempty"`
	SharedDnnConfigurationsId       string                             `json:"sharedDnnConfigurationsId,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	ExpectedUeBehavioursList        map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`
	SuggestedPacketNumDlList        map[string]SuggestedPacketNumDl    `json:"suggestedPacketNumDlList,omitempty"`
	SingleNssai                     Snssai                             `json:"singleNssai"`
	DnnConfigurations               map[string]DnnConfiguration        `json:"dnnConfigurations,omitempty"`
	ThreeGppChargingCharacteristics string                             `json:"3gppChargingCharacteristics,omitempty"`
}
