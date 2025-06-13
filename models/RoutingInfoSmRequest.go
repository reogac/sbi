/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RoutingInfoSmRequest struct {
	CorrelationId     string `json:"correlationId,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	IpSmGwInd         *bool  `json:"ipSmGwInd,omitempty"`
}
