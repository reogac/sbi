/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:26:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             int             `json:"status"`
	Detail             string          `json:"detail,omitempty"`
	Cause              string          `json:"cause,omitempty"`
}
