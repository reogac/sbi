/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ContextDataSets struct {
	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
	Amf3Gpp                       *Amf3GppAccessRegistration      `json:"amf3Gpp,omitempty"`
	AmfNon3Gpp                    *AmfNon3GppAccessRegistration   `json:"amfNon3Gpp,omitempty"`
	SdmSubscriptions              []SdmSubscription               `json:"sdmSubscriptions,omitempty"`
	SmsfNon3GppAccess             *SmsfRegistration               `json:"smsfNon3GppAccess,omitempty"`
	RoamingInfo                   *RoamingInfoUpdate              `json:"roamingInfo,omitempty"`
	PeiInfo                       *PeiUpdateInfo                  `json:"peiInfo,omitempty"`
	EeSubscriptions               []EeSubscription                `json:"eeSubscriptions,omitempty"`
	Smsf3GppAccess                *SmsfRegistration               `json:"smsf3GppAccess,omitempty"`
	SmfRegistrations              []SmfRegistration               `json:"smfRegistrations,omitempty"`
	IpSmGw                        *IpSmGwRegistration             `json:"ipSmGw,omitempty"`
}
