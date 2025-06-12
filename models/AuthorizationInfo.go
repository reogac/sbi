/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizationInfo struct {
	AfId                   string       `json:"afId,omitempty"`
	NefId                  string       `json:"nefId,omitempty"`
	ValidityTime           string       `json:"validityTime,omitempty"`
	ContextInfo            *ContextInfo `json:"contextInfo,omitempty"`
	Snssai                 Snssai       `json:"snssai"`
	Dnn                    string       `json:"dnn"`
	MtcProviderInformation string       `json:"mtcProviderInformation"`
	AuthUpdateCallbackUri  string       `json:"authUpdateCallbackUri"`
}
