/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:35 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NsiInformation struct {
	NrfNfMgtUri       string          `json:"nrfNfMgtUri,omitempty"`
	NrfAccessTokenUri string          `json:"nrfAccessTokenUri,omitempty"`
	NrfOauth2Required map[string]bool `json:"nrfOauth2Required,omitempty"`
	NrfId             string          `json:"nrfId"`
	NsiId             string          `json:"nsiId,omitempty"`
}
