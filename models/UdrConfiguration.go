/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:25 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UdrConfiguration struct {
	SmSub   string `json:"smSub,omitempty"`
	SmfSel  string `json:"smfSel,omitempty"`
	Url     string `json:"url"`
	Mock    *bool  `json:"mock,omitempty"`
	DbName  string `json:"dbName"`
	AuthSub string `json:"authSub"`
	AmSub   string `json:"amSub,omitempty"`
}
