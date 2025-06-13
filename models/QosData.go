/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	QosId                string `json:"qosId"`
	Qnc                  *bool  `json:"qnc,omitempty"`
}
