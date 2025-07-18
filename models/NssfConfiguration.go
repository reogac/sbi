/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NssfConfiguration struct {
	AmfSets   []AmfSetConfiguration   `json:"amfSets"`
	Slices    []SliceConfiguration    `json:"slices,omitempty"`
	PlmnPeers []HomePlmnConfiguration `json:"plmnPeers,omitempty"`
}
