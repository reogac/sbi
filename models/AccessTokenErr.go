/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 17:39:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenErr struct {
	ErrorUri         string `json:"error_uri,omitempty"`
	Error            Error  `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
}
