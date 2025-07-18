/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UdrConfiguration struct {
	Mock    *bool  `json:"mock,omitempty"`
	DbName  string `json:"dbName"`
	AuthSub string `json:"authSub"`
	AmSub   string `json:"amSub,omitempty"`
	SmSub   string `json:"smSub,omitempty"`
	SmfSel  string `json:"smfSel,omitempty"`
	Url     string `json:"url"`
}
