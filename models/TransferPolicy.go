/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferPolicy struct {
	RecTimeInt    TimeWindow `json:"recTimeInt"`
	TransPolicyId int        `json:"transPolicyId"`
	MaxBitRateDl  string     `json:"maxBitRateDl,omitempty"`
	MaxBitRateUl  string     `json:"maxBitRateUl,omitempty"`
	RatingGroup   int        `json:"ratingGroup"`
}
