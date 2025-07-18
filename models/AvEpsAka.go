/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEpsAka struct {
	Autn   string    `json:"autn"`
	Kasme  string    `json:"kasme"`
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
}
