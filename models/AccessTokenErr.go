/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:47 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenErr struct {
	Error            Error  `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorUri         string `json:"error_uri,omitempty"`
}
