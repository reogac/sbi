/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsSessPolCtrlData struct {
	FiveQis           []int  `json:"5qis,omitempty"`
	MaxMbsArpLevel    *int   `json:"maxMbsArpLevel,omitempty"`
	MaxMbsSessionAmbr string `json:"maxMbsSessionAmbr,omitempty"`
	MaxGbr            string `json:"maxGbr,omitempty"`
	SuppFeat          string `json:"suppFeat,omitempty"`
}
