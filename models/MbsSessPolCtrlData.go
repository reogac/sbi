/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsSessPolCtrlData struct {
	MaxGbr            string `json:"maxGbr,omitempty"`
	SuppFeat          string `json:"suppFeat,omitempty"`
	FiveQis           []int  `json:"5qis,omitempty"`
	MaxMbsArpLevel    *int   `json:"maxMbsArpLevel,omitempty"`
	MaxMbsSessionAmbr string `json:"maxMbsSessionAmbr,omitempty"`
}
