/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NssfConfiguration struct {
	Slices    []SliceConfiguration    `json:"slices,omitempty"`
	PlmnPeers []HomePlmnConfiguration `json:"plmnPeers,omitempty"`
	AmfSets   []AmfSetConfiguration   `json:"amfSets"`
}
