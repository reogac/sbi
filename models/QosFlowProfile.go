/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowProfile struct {
	FiveQi                int                    `json:"5qi"`
	Dynamic5Qi            *Dynamic5Qi            `json:"dynamic5Qi,omitempty"`
	GbrQosFlowInfo        *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	AdditionalQosFlowInfo AdditionalQosFlowInfo  `json:"additionalQosFlowInfo,omitempty"`
	QosRepPeriod          *int                   `json:"qosRepPeriod,omitempty"`
	NonDynamic5Qi         *NonDynamic5Qi         `json:"nonDynamic5Qi,omitempty"`
	Arp                   *Arp                   `json:"arp,omitempty"`
	Rqa                   ReflectiveQoSAttribute `json:"rqa,omitempty"`
	QosMonitoringReq      QosMonitoringReq       `json:"qosMonitoringReq,omitempty"`
}
