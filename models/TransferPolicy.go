/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferPolicy struct {
	RatingGroup   int        `json:"ratingGroup"`
	RecTimeInt    TimeWindow `json:"recTimeInt"`
	TransPolicyId int        `json:"transPolicyId"`
	MaxBitRateDl  string     `json:"maxBitRateDl,omitempty"`
	MaxBitRateUl  string     `json:"maxBitRateUl,omitempty"`
}
