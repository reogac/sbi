/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEapAkaPrime struct {
	Rand    string `json:"rand"`
	Xres    string `json:"xres"`
	Autn    string `json:"autn"`
	CkPrime string `json:"ckPrime"`
	IkPrime string `json:"ikPrime"`
	AvType  AvType `json:"avType"`
}
