/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CongestionInfo struct {
	TopAppListDl []TopApplication `json:"topAppListDl,omitempty"`
	CongType     CongestionType   `json:"congType"`
	TimeIntev    TimeWindow       `json:"timeIntev"`
	Nsi          ThresholdLevel   `json:"nsi"`
	Confidence   *int             `json:"confidence,omitempty"`
	TopAppListUl []TopApplication `json:"topAppListUl,omitempty"`
}
