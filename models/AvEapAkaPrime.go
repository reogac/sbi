/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:38 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEapAkaPrime struct {
	AvType  AvType `json:"avType"`
	Rand    string `json:"rand"`
	Xres    string `json:"xres"`
	Autn    string `json:"autn"`
	CkPrime string `json:"ckPrime"`
	IkPrime string `json:"ikPrime"`
}
