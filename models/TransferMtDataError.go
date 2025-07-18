/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMtDataError struct {
	Cause              string          `json:"cause,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	RemoteError        *bool           `json:"remoteError,omitempty"`
	Title              string          `json:"title,omitempty"`
	Instance           string          `json:"instance,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty"`
	Type               string          `json:"type,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty"`
	MaxWaitingTime     *int            `json:"maxWaitingTime,omitempty"`
	Detail             string          `json:"detail,omitempty"`
	NrfId              string          `json:"nrfId,omitempty"`
	Status             int             `json:"status"`
}
