/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SupportedPlmnItem struct {
	Slices []Snssai `json:"slices,omitempty"`
	PlmnId PlmnId   `json:"plmnId"`
}
