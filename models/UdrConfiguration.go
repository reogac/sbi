/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Apr 29 09:34:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UdrConfiguration struct {
	SmfSel  string `json:"smfSel,omitempty"`
	Url     string `json:"url"`
	Mock    *bool  `json:"mock,omitempty"`
	DbName  string `json:"dbName"`
	AuthSub string `json:"authSub"`
	AmSub   string `json:"amSub,omitempty"`
	SmSub   string `json:"smSub,omitempty"`
}
