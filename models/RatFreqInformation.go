/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RatFreqInformation struct {
	AllFreq         *bool             `json:"allFreq,omitempty"`
	AllRat          *bool             `json:"allRat,omitempty"`
	Freq            *int              `json:"freq,omitempty"`
	RatType         RatType           `json:"ratType,omitempty"`
	SvcExpThreshold *ThresholdLevel   `json:"svcExpThreshold,omitempty"`
	MatchingDir     MatchingDirection `json:"matchingDir,omitempty"`
}
