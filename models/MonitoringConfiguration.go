/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
}
