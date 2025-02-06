/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Feb  6 13:48:14 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	Type               string          `json:"type,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             int             `json:"status"`
	Cause              string          `json:"cause,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
}
