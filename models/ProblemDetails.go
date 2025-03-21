/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 09:41:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	Instance           string          `json:"instance,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             int             `json:"status"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
}
