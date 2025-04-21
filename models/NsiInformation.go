/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Apr 21 15:01:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NsiInformation struct {
	NrfId             string          `json:"nrfId"`
	NsiId             string          `json:"nsiId,omitempty"`
	NrfNfMgtUri       string          `json:"nrfNfMgtUri,omitempty"`
	NrfAccessTokenUri string          `json:"nrfAccessTokenUri,omitempty"`
	NrfOauth2Required map[string]bool `json:"nrfOauth2Required,omitempty"`
}
