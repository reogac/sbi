/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistrationModification struct {
	SmfInstanceId string `json:"smfInstanceId"`
	SmfSetId      string `json:"smfSetId,omitempty"`
	PgwFqdn       string `json:"pgwFqdn,omitempty"`
}
