/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VplmnQos struct {
	SessionAmbr *Ambr  `json:"sessionAmbr,omitempty"`
	MaxFbrDl    string `json:"maxFbrDl,omitempty"`
	MaxFbrUl    string `json:"maxFbrUl,omitempty"`
	GuaFbrDl    string `json:"guaFbrDl,omitempty"`
	GuaFbrUl    string `json:"guaFbrUl,omitempty"`
	FiveQi      *int   `json:"5qi,omitempty"`
	Arp         *Arp   `json:"arp,omitempty"`
}
