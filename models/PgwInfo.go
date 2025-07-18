/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PgwInfo struct {
	EpdgInd          *bool      `json:"epdgInd,omitempty"`
	PcfId            string     `json:"pcfId,omitempty"`
	RegistrationTime string     `json:"registrationTime,omitempty"`
	Dnn              string     `json:"dnn"`
	PgwFqdn          string     `json:"pgwFqdn"`
	PgwIpAddr        *IpAddress `json:"pgwIpAddr,omitempty"`
	PlmnId           *PlmnId    `json:"plmnId,omitempty"`
}
