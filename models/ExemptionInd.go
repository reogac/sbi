/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExemptionInd struct {
	SnssaiDnnCongestion  *bool `json:"snssaiDnnCongestion,omitempty"`
	DnnCongestion        *bool `json:"dnnCongestion,omitempty"`
	SnssaiOnlyCongestion *bool `json:"snssaiOnlyCongestion,omitempty"`
}
