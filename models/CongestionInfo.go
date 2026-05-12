/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CongestionInfo struct {
	TimeIntev    TimeWindow       `json:"timeIntev"`
	Nsi          ThresholdLevel   `json:"nsi"`
	Confidence   *int             `json:"confidence,omitempty"`
	TopAppListUl []TopApplication `json:"topAppListUl,omitempty"`
	TopAppListDl []TopApplication `json:"topAppListDl,omitempty"`
	CongType     CongestionType   `json:"congType"`
}
