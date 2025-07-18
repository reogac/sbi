/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMtDataError struct {
	Detail             string          `json:"detail,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	Type               string          `json:"type,omitempty"`
	Title              string          `json:"title,omitempty"`
	Status             int             `json:"status"`
	NrfId              string          `json:"nrfId,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Instance           string          `json:"instance,omitempty"`
}
