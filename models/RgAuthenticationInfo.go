/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RgAuthenticationInfo struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	Suci              string `json:"suci"`
	AuthenticatedInd  bool   `json:"authenticatedInd"`
}
