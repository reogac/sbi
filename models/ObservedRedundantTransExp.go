/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ObservedRedundantTransExp struct {
	VarPktDelayUl    *float64 `json:"varPktDelayUl,omitempty"`
	AvgPktDelayDl    *int     `json:"avgPktDelayDl,omitempty"`
	VarPktDelayDl    *float64 `json:"varPktDelayDl,omitempty"`
	AvgPktDropRateUl *int     `json:"avgPktDropRateUl,omitempty"`
	VarPktDropRateUl *float64 `json:"varPktDropRateUl,omitempty"`
	AvgPktDropRateDl *int     `json:"avgPktDropRateDl,omitempty"`
	VarPktDropRateDl *float64 `json:"varPktDropRateDl,omitempty"`
	AvgPktDelayUl    *int     `json:"avgPktDelayUl,omitempty"`
}
