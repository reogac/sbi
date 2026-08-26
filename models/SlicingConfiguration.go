/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:59:14 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SlicingConfiguration struct {
	AmfSets   []AmfSetConfiguration   `json:"amfSets"`
	Slices    []SliceConfiguration    `json:"slices,omitempty"`
	PlmnPeers []HomePlmnConfiguration `json:"plmnPeers,omitempty"`
}
