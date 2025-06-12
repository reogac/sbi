/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PgwInfo struct {
	PlmnId           *PlmnId    `json:"plmnId,omitempty"`
	EpdgInd          *bool      `json:"epdgInd,omitempty"`
	PcfId            string     `json:"pcfId,omitempty"`
	RegistrationTime string     `json:"registrationTime,omitempty"`
	Dnn              string     `json:"dnn"`
	PgwFqdn          string     `json:"pgwFqdn"`
	PgwIpAddr        *IpAddress `json:"pgwIpAddr,omitempty"`
}
