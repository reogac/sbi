/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:26 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	AuthenticationVector *AuthenticationVector              `json:"authenticationVector,omitempty"`
	RoutingId            string                             `json:"routingId,omitempty"`
	PvsInfo              []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	SupportedFeatures    string                             `json:"supportedFeatures,omitempty"`
	AuthType             AuthType                           `json:"authType"`
	Supi                 string                             `json:"supi,omitempty"`
	AkmaInd              *bool                              `json:"akmaInd,omitempty"`
	AuthAaa              *bool                              `json:"authAaa,omitempty"`
	AmData               *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
