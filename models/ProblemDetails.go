/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProblemDetails struct {
	Status             int             `json:"status"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Type               string          `json:"type,omitempty"`
	Title              string          `json:"title,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
}
