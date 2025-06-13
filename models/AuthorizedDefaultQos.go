/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizedDefaultQos struct {
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	Arp                *Arp   `json:"arp,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
}
