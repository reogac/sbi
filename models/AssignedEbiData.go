/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignedEbiData struct {
	PduSessionId    int             `json:"pduSessionId"`
	AssignedEbiList []EbiArpMapping `json:"assignedEbiList"`
	FailedArpList   []Arp           `json:"failedArpList,omitempty"`
	ReleasedEbiList []int           `json:"releasedEbiList,omitempty"`
	ModifiedEbiList []int           `json:"modifiedEbiList,omitempty"`
}
