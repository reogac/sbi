/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:59:14 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSetConfiguration struct {
	SetId  string   `json:"setId"`
	Region string   `json:"region"`
	Slices []Snssai `json:"slices,omitempty"`
}
