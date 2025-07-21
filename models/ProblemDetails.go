/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Jul 21 09:39:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	Title              string          `json:"title,omitempty"`
	Status             int             `json:"status"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Type               string          `json:"type,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
}
