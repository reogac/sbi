/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:27 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RoutingInfoSmRequest struct {
	IpSmGwInd         *bool  `json:"ipSmGwInd,omitempty"`
	CorrelationId     string `json:"correlationId,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
