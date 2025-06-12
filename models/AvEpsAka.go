/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEpsAka struct {
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
	Autn   string    `json:"autn"`
	Kasme  string    `json:"kasme"`
	AvType HssAvType `json:"avType"`
}
