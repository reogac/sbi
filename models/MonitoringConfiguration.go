/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
}
