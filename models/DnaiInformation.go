/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnaiInformation struct {
	NoDnaiChangeInd     *bool  `json:"noDnaiChangeInd,omitempty"`
	NoLocalPsaChangeInd *bool  `json:"noLocalPsaChangeInd,omitempty"`
	Dnai                string `json:"dnai"`
}
