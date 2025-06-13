/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
