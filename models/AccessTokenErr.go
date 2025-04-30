/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 17:37:55 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenErr struct {
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorUri         string `json:"error_uri,omitempty"`
	Error            Error  `json:"error"`
}
