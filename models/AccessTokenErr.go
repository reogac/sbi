/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenErr struct {
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorUri         string `json:"error_uri,omitempty"`
	Error            Error  `json:"error"`
}
