/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:38:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenErr struct {
	Error            Error  `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorUri         string `json:"error_uri,omitempty"`
}
