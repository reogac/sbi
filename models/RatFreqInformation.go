/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RatFreqInformation struct {
	AllRat          *bool             `json:"allRat,omitempty"`
	Freq            *int              `json:"freq,omitempty"`
	RatType         RatType           `json:"ratType,omitempty"`
	SvcExpThreshold *ThresholdLevel   `json:"svcExpThreshold,omitempty"`
	MatchingDir     MatchingDirection `json:"matchingDir,omitempty"`
	AllFreq         *bool             `json:"allFreq,omitempty"`
}
