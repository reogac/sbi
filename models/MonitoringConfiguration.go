/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
}
