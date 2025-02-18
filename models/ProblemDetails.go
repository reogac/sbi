/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:05:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models
type ProblemDetails struct {
	 InvalidParams	[]InvalidParam	`json:"invalidParams,omitempty"`
	 AccessTokenError	*AccessTokenErr	`json:"accessTokenError,omitempty"`
	 NrfId	string	`json:"nrfId,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 Status	int	`json:"status"`
	 Detail	string	`json:"detail,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 AccessTokenRequest	*AccessTokenReq	`json:"accessTokenRequest,omitempty"`
	 Type	string	`json:"type,omitempty"`
	 Instance	string	`json:"instance,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
}
