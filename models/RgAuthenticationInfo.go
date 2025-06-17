/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:57 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RgAuthenticationInfo struct {
	Suci              string `json:"suci"`
	AuthenticatedInd  bool   `json:"authenticatedInd"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
