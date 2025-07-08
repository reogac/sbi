/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NumberAverage struct {
	Number   float64  `json:"number"`
	Variance float64  `json:"variance"`
	Skewness *float64 `json:"skewness,omitempty"`
}
