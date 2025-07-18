/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizedDefaultQos struct {
	Arp                *Arp   `json:"arp,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
}
