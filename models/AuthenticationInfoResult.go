/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	SupportedFeatures    string                             `json:"supportedFeatures,omitempty"`
	AkmaInd              *bool                              `json:"akmaInd,omitempty"`
	PvsInfo              []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData               *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthType             AuthType                           `json:"authType"`
	AuthAaa              *bool                              `json:"authAaa,omitempty"`
	RoutingId            string                             `json:"routingId,omitempty"`
	AuthenticationVector *AuthenticationVector              `json:"authenticationVector,omitempty"`
	Supi                 string                             `json:"supi,omitempty"`
}
