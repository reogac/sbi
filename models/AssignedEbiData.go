/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignedEbiData struct {
	ReleasedEbiList []int           `json:"releasedEbiList,omitempty"`
	ModifiedEbiList []int           `json:"modifiedEbiList,omitempty"`
	PduSessionId    int             `json:"pduSessionId"`
	AssignedEbiList []EbiArpMapping `json:"assignedEbiList"`
	FailedArpList   []Arp           `json:"failedArpList,omitempty"`
}
