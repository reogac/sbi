/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenErr struct {
	Error            Error  `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorUri         string `json:"error_uri,omitempty"`
}
