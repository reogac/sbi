/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:26 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvImsGbaEapAka struct {
	Xres   string    `json:"xres"`
	Autn   string    `json:"autn"`
	Ck     string    `json:"ck"`
	Ik     string    `json:"ik"`
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
}
