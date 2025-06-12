/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	SupportedFeatures    string                             `json:"supportedFeatures,omitempty"`
	AkmaInd              *bool                              `json:"akmaInd,omitempty"`
	AuthAaa              *bool                              `json:"authAaa,omitempty"`
	PvsInfo              []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData               *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthType             AuthType                           `json:"authType"`
	AuthenticationVector *AuthenticationVector              `json:"authenticationVector,omitempty"`
	Supi                 string                             `json:"supi,omitempty"`
	RoutingId            string                             `json:"routingId,omitempty"`
}
