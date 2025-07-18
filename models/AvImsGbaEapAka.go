/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
