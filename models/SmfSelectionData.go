/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:43 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfSelectionData struct {
	UnsuppDnn     *bool                              `json:"unsuppDnn,omitempty"`
	Candidates    map[string]CandidateForReplacement `json:"candidates,omitempty"`
	Snssai        *Snssai                            `json:"snssai,omitempty"`
	MappingSnssai *Snssai                            `json:"mappingSnssai,omitempty"`
	Dnn           string                             `json:"dnn,omitempty"`
}
