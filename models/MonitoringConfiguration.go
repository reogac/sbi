/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
}
