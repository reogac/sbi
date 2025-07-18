/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	QosId                string `json:"qosId"`
}
