/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	AuthType             AuthType                           `json:"authType"`
	SupportedFeatures    string                             `json:"supportedFeatures,omitempty"`
	RoutingId            string                             `json:"routingId,omitempty"`
	PvsInfo              []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData               *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthenticationVector *AuthenticationVector              `json:"authenticationVector,omitempty"`
	Supi                 string                             `json:"supi,omitempty"`
	AkmaInd              *bool                              `json:"akmaInd,omitempty"`
	AuthAaa              *bool                              `json:"authAaa,omitempty"`
}
