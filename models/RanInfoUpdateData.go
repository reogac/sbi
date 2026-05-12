/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:29 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanInfoUpdateData struct {
	Id              string            `json:"id"`
	SupportedTAList []SupportedTAItem `json:"supportedTAList,omitempty"`
}
