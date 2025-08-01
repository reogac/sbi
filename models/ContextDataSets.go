/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ContextDataSets struct {
	EeSubscriptions               []EeSubscription                `json:"eeSubscriptions,omitempty"`
	Smsf3GppAccess                *SmsfRegistration               `json:"smsf3GppAccess,omitempty"`
	SmsfNon3GppAccess             *SmsfRegistration               `json:"smsfNon3GppAccess,omitempty"`
	SmfRegistrations              []SmfRegistration               `json:"smfRegistrations,omitempty"`
	PeiInfo                       *PeiUpdateInfo                  `json:"peiInfo,omitempty"`
	Amf3Gpp                       *Amf3GppAccessRegistration      `json:"amf3Gpp,omitempty"`
	AmfNon3Gpp                    *AmfNon3GppAccessRegistration   `json:"amfNon3Gpp,omitempty"`
	IpSmGw                        *IpSmGwRegistration             `json:"ipSmGw,omitempty"`
	RoamingInfo                   *RoamingInfoUpdate              `json:"roamingInfo,omitempty"`
	SdmSubscriptions              []SdmSubscription               `json:"sdmSubscriptions,omitempty"`
	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
}
