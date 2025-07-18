/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfSelectionData struct {
	MappingSnssai *Snssai                            `json:"mappingSnssai,omitempty"`
	Dnn           string                             `json:"dnn,omitempty"`
	UnsuppDnn     *bool                              `json:"unsuppDnn,omitempty"`
	Candidates    map[string]CandidateForReplacement `json:"candidates,omitempty"`
	Snssai        *Snssai                            `json:"snssai,omitempty"`
}
