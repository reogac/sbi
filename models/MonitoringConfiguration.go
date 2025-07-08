/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
}
