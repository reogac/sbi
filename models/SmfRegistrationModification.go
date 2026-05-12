/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:39 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistrationModification struct {
	SmfInstanceId string `json:"smfInstanceId"`
	SmfSetId      string `json:"smfSetId,omitempty"`
	PgwFqdn       string `json:"pgwFqdn,omitempty"`
}
