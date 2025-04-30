/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VplmnQos struct {
	MaxFbrDl    string `json:"maxFbrDl,omitempty"`
	MaxFbrUl    string `json:"maxFbrUl,omitempty"`
	GuaFbrDl    string `json:"guaFbrDl,omitempty"`
	GuaFbrUl    string `json:"guaFbrUl,omitempty"`
	FiveQi      *int   `json:"5qi,omitempty"`
	Arp         *Arp   `json:"arp,omitempty"`
	SessionAmbr *Ambr  `json:"sessionAmbr,omitempty"`
}
