/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEapAkaPrime struct {
	Autn    string `json:"autn"`
	CkPrime string `json:"ckPrime"`
	IkPrime string `json:"ikPrime"`
	AvType  AvType `json:"avType"`
	Rand    string `json:"rand"`
	Xres    string `json:"xres"`
}
