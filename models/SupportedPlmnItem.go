/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:38:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SupportedPlmnItem struct {
	PlmnId PlmnId   `json:"plmnId"`
	Slices []Snssai `json:"slices,omitempty"`
}
