/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	AuthenticationVector *AuthenticationVector              `json:"authenticationVector,omitempty"`
	Supi                 string                             `json:"supi,omitempty"`
	AkmaInd              *bool                              `json:"akmaInd,omitempty"`
	RoutingId            string                             `json:"routingId,omitempty"`
	PvsInfo              []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData               *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthType             AuthType                           `json:"authType"`
	SupportedFeatures    string                             `json:"supportedFeatures,omitempty"`
	AuthAaa              *bool                              `json:"authAaa,omitempty"`
}
