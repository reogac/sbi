/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:37 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IdTranslationResult struct {
	Gpsi              string   `json:"gpsi,omitempty"`
	AdditionalSupis   []string `json:"additionalSupis,omitempty"`
	AdditionalGpsis   []string `json:"additionalGpsis,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
	Supi              string   `json:"supi"`
}
