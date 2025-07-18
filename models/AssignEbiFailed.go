/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignEbiFailed struct {
	PduSessionId  int   `json:"pduSessionId"`
	FailedArpList []Arp `json:"failedArpList,omitempty"`
}
