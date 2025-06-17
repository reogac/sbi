/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RatFreqInformation struct {
	MatchingDir     MatchingDirection `json:"matchingDir,omitempty"`
	AllFreq         *bool             `json:"allFreq,omitempty"`
	AllRat          *bool             `json:"allRat,omitempty"`
	Freq            *int              `json:"freq,omitempty"`
	RatType         RatType           `json:"ratType,omitempty"`
	SvcExpThreshold *ThresholdLevel   `json:"svcExpThreshold,omitempty"`
}
