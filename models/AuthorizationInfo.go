/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizationInfo struct {
	Snssai                 Snssai       `json:"snssai"`
	Dnn                    string       `json:"dnn"`
	MtcProviderInformation string       `json:"mtcProviderInformation"`
	AuthUpdateCallbackUri  string       `json:"authUpdateCallbackUri"`
	AfId                   string       `json:"afId,omitempty"`
	NefId                  string       `json:"nefId,omitempty"`
	ValidityTime           string       `json:"validityTime,omitempty"`
	ContextInfo            *ContextInfo `json:"contextInfo,omitempty"`
}
