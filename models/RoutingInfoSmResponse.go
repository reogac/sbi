/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:39 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RoutingInfoSmResponse struct {
	Supi        string            `json:"supi,omitempty"`
	Smsf3Gpp    *SmsfRegistration `json:"smsf3Gpp,omitempty"`
	SmsfNon3Gpp *SmsfRegistration `json:"smsfNon3Gpp,omitempty"`
	IpSmGw      *IpSmGwInfo       `json:"ipSmGw,omitempty"`
	SmsRouter   *SmsRouterInfo    `json:"smsRouter,omitempty"`
}
