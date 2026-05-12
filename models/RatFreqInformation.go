/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RatFreqInformation struct {
	Freq            *int              `json:"freq,omitempty"`
	RatType         RatType           `json:"ratType,omitempty"`
	SvcExpThreshold *ThresholdLevel   `json:"svcExpThreshold,omitempty"`
	MatchingDir     MatchingDirection `json:"matchingDir,omitempty"`
	AllFreq         *bool             `json:"allFreq,omitempty"`
	AllRat          *bool             `json:"allRat,omitempty"`
}
