/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Mar 19 09:38:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	Title              string          `json:"title,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	Status             int             `json:"status"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
}
