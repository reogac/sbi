/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:58 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	FiveQi               *int   `json:"5qi,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
}
