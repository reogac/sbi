/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SupportedTAItem struct {
	Tac    string   `json:"tac"`
	Slices []Snssai `json:"slices,omitempty"`
}
