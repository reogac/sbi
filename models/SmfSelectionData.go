/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
