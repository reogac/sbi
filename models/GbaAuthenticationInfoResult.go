/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:54 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GbaAuthenticationInfoResult struct {
	SupportedFeatures string       `json:"supportedFeatures,omitempty"`
	ThreeGAkaAv       *ThreeGAkaAv `json:"3gAkaAv,omitempty"`
}
