/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	QosId                string `json:"qosId"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
}
