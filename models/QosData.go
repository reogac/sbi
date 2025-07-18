/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
}
