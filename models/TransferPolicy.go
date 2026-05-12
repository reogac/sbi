/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferPolicy struct {
	MaxBitRateUl  string     `json:"maxBitRateUl,omitempty"`
	RatingGroup   int        `json:"ratingGroup"`
	RecTimeInt    TimeWindow `json:"recTimeInt"`
	TransPolicyId int        `json:"transPolicyId"`
	MaxBitRateDl  string     `json:"maxBitRateDl,omitempty"`
}
