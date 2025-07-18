/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MonitoringConfiguration struct {
	LocationReportingConfiguration *LocationReportingConfiguration   `json:"locationReportingConfiguration,omitempty"`
	AssociationType                AssociationType                   `json:"associationType,omitempty"`
	EventType                      EventType                         `json:"eventType"`
	MaximumResponseTime            *int                              `json:"maximumResponseTime,omitempty"`
	SingleNssai                    *Snssai                           `json:"singleNssai,omitempty"`
	MtcProviderInformation         string                            `json:"mtcProviderInformation,omitempty"`
	AfId                           string                            `json:"afId,omitempty"`
	IdleStatusInd                  *bool                             `json:"idleStatusInd,omitempty"`
	MaximumLatency                 *int                              `json:"maximumLatency,omitempty"`
	SuggestedPacketNumDl           *int                              `json:"suggestedPacketNumDl,omitempty"`
	Dnn                            string                            `json:"dnn,omitempty"`
	PduSessionStatusCfg            *PduSessionStatusCfg              `json:"pduSessionStatusCfg,omitempty"`
	ReachabilityForSmsCfg          ReachabilityForSmsConfiguration   `json:"reachabilityForSmsCfg,omitempty"`
	ReachabilityForDataCfg         *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	ImmediateFlag                  *bool                             `json:"immediateFlag,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration   `json:"datalinkReportCfg,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg              `json:"lossConnectivityCfg,omitempty"`
}
