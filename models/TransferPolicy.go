/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferPolicy struct {
	MaxBitRateDl  string     `json:"maxBitRateDl,omitempty"`
	MaxBitRateUl  string     `json:"maxBitRateUl,omitempty"`
	RatingGroup   int        `json:"ratingGroup"`
	RecTimeInt    TimeWindow `json:"recTimeInt"`
	TransPolicyId int        `json:"transPolicyId"`
}
