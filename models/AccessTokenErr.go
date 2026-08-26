/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenErr struct {
	ErrorUri         string `json:"error_uri,omitempty"`
	Error            Error  `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
}
