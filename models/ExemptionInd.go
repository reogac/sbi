/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExemptionInd struct {
	DnnCongestion        *bool `json:"dnnCongestion,omitempty"`
	SnssaiOnlyCongestion *bool `json:"snssaiOnlyCongestion,omitempty"`
	SnssaiDnnCongestion  *bool `json:"snssaiDnnCongestion,omitempty"`
}
