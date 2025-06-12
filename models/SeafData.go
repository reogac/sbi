/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SeafData struct {
	NgKsi                NgKsi  `json:"ngKsi"`
	KeyAmf               KeyAmf `json:"keyAmf"`
	Nh                   string `json:"nh,omitempty"`
	Ncc                  *int   `json:"ncc,omitempty"`
	KeyAmfChangeInd      *bool  `json:"keyAmfChangeInd,omitempty"`
	KeyAmfHDerivationInd *bool  `json:"keyAmfHDerivationInd,omitempty"`
}
