/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 15:35:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NssfConfiguration struct {
	AmfSets   []AmfSetConfiguration   `json:"amfSets"`
	Slices    []SliceConfiguration    `json:"slices,omitempty"`
	PlmnPeers []HomePlmnConfiguration `json:"plmnPeers,omitempty"`
}
