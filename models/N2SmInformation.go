/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInformation struct {
	PduSessionId   int            `json:"pduSessionId"`
	N2InfoContent  *N2InfoContent `json:"n2InfoContent,omitempty"`
	SNssai         *Snssai        `json:"sNssai,omitempty"`
	HomePlmnSnssai *Snssai        `json:"homePlmnSnssai,omitempty"`
	IwkSnssai      *Snssai        `json:"iwkSnssai,omitempty"`
	SubjectToHo    *bool          `json:"subjectToHo,omitempty"`
}
