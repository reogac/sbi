/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
