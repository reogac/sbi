/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GbaAuthenticationInfoRequest struct {
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	AuthType              GbaAuthType            `json:"authType"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
}
