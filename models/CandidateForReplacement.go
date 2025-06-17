/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:00 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CandidateForReplacement struct {
	Snssai Snssai   `json:"snssai"`
	Dnns   []string `json:"dnns,omitempty"`
}
