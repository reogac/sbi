/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	QosId                string `json:"qosId"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
}
