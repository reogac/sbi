/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IdTranslationResult struct {
	AdditionalSupis   []string `json:"additionalSupis,omitempty"`
	AdditionalGpsis   []string `json:"additionalGpsis,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
	Supi              string   `json:"supi"`
	Gpsi              string   `json:"gpsi,omitempty"`
}
