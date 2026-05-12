/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:25 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NssfConfiguration struct {
	AmfSets   []AmfSetConfiguration   `json:"amfSets"`
	Slices    []SliceConfiguration    `json:"slices,omitempty"`
	PlmnPeers []HomePlmnConfiguration `json:"plmnPeers,omitempty"`
}
