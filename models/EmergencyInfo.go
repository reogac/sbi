/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EmergencyInfo struct {
	EpdgInd       *bool      `json:"epdgInd,omitempty"`
	PlmnId        *PlmnId    `json:"plmnId,omitempty"`
	PgwFqdn       string     `json:"pgwFqdn,omitempty"`
	PgwIpAddress  *IpAddress `json:"pgwIpAddress,omitempty"`
	SmfInstanceId string     `json:"smfInstanceId,omitempty"`
}
