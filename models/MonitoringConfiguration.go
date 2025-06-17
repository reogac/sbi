/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
}
