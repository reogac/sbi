/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
}
