/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SupportedTAItem struct {
	Tac      string              `json:"tac"`
	PlmnList []SupportedPlmnItem `json:"plmnList,omitempty"`
}
