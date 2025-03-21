/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	Type               string          `json:"type,omitempty"`
	Status             int             `json:"status"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Title              string          `json:"title,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
}
