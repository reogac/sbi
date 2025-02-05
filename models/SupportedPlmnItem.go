/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Feb  5 09:54:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SupportedPlmnItem struct {
	PlmnId PlmnId   `json:"plmnId"`
	Slices []Snssai `json:"slices"`
}
