/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizedDefaultQos struct {
	GbrUl              string `json:"gbrUl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	Arp                *Arp   `json:"arp,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
}
