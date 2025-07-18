/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SeafData struct {
	KeyAmf               KeyAmf `json:"keyAmf"`
	Nh                   string `json:"nh,omitempty"`
	Ncc                  *int   `json:"ncc,omitempty"`
	KeyAmfChangeInd      *bool  `json:"keyAmfChangeInd,omitempty"`
	KeyAmfHDerivationInd *bool  `json:"keyAmfHDerivationInd,omitempty"`
	NgKsi                NgKsi  `json:"ngKsi"`
}
