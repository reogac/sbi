/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
