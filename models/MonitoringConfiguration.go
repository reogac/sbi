/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
}
