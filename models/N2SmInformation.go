/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
