/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	QosId                string `json:"qosId"`
	GbrDl                string `json:"gbrDl,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
}
