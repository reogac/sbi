/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Feb  5 17:39:08 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SupportedTAItem struct {
	Tac    string   `json:"tac"`
	PlmnId PlmnId   `json:"plmnId"`
	Slices []Snssai `json:"slices,omitempty"`
}
