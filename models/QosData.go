/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	Qnc                  *bool  `json:"qnc,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	QosId                string `json:"qosId"`
	FiveQi               *int   `json:"5qi,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
}
