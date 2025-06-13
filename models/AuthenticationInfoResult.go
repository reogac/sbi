/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	PvsInfo              []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AuthType             AuthType                           `json:"authType"`
	Supi                 string                             `json:"supi,omitempty"`
	AkmaInd              *bool                              `json:"akmaInd,omitempty"`
	AuthAaa              *bool                              `json:"authAaa,omitempty"`
	RoutingId            string                             `json:"routingId,omitempty"`
	AmData               *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SupportedFeatures    string                             `json:"supportedFeatures,omitempty"`
	AuthenticationVector *AuthenticationVector              `json:"authenticationVector,omitempty"`
}
