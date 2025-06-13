/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMtDataError struct {
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Status             int             `json:"status"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	Title              string          `json:"title,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	Cause              string          `json:"cause,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Type               string          `json:"type,omitempty"`
}
