/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:32 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceAreaId struct {
	PlmnId PlmnId `json:"plmnId"`
	Lac    string `json:"lac"`
	Sac    string `json:"sac"`
}
