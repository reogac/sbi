/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Type               string          `json:"type,omitempty"`
	Title              string          `json:"title,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Status             int             `json:"status"`
	Detail             string          `json:"detail,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
}
