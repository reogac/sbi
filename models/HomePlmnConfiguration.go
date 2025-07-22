/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 15:13:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HomePlmnConfiguration struct {
	Id     PlmnId            `json:"id"`
	Slices []MappingOfSnssai `json:"slices,omitempty"`
	Sepps  []string          `json:"sepps,omitempty"`
}
