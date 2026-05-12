/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:38 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	AuthAaa              *bool                              `json:"authAaa,omitempty"`
	RoutingId            string                             `json:"routingId,omitempty"`
	PvsInfo              []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData               *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SupportedFeatures    string                             `json:"supportedFeatures,omitempty"`
	AuthenticationVector *AuthenticationVector              `json:"authenticationVector,omitempty"`
	Supi                 string                             `json:"supi,omitempty"`
	AkmaInd              *bool                              `json:"akmaInd,omitempty"`
	AuthType             AuthType                           `json:"authType"`
}
