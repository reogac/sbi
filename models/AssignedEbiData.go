/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignedEbiData struct {
	AssignedEbiList []EbiArpMapping `json:"assignedEbiList"`
	FailedArpList   []Arp           `json:"failedArpList,omitempty"`
	ReleasedEbiList []int           `json:"releasedEbiList,omitempty"`
	ModifiedEbiList []int           `json:"modifiedEbiList,omitempty"`
	PduSessionId    int             `json:"pduSessionId"`
}
