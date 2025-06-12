/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	Dnn                            string                            `json:"dnn,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	EventType                      EventType                         `json:"eventType"`
}
