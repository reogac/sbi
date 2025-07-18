/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnnConfiguration struct {
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
}
