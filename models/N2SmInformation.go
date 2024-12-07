/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInformation struct {
	SubjectToHo    *bool          `json:"subjectToHo,omitempty"`
	PduSessionId   int            `json:"pduSessionId"`
	N2InfoContent  *N2InfoContent `json:"n2InfoContent,omitempty"`
	SNssai         *Snssai        `json:"sNssai,omitempty"`
	HomePlmnSnssai *Snssai        `json:"homePlmnSnssai,omitempty"`
	IwkSnssai      *Snssai        `json:"iwkSnssai,omitempty"`
}
