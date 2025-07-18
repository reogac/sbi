/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssAuthenticationInfoResult struct {
	SupportedFeatures        string                   `json:"supportedFeatures,omitempty"`
	HssAuthenticationVectors HssAuthenticationVectors `json:"hssAuthenticationVectors"`
}
