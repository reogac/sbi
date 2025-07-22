/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
}
