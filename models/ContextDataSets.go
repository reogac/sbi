/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ContextDataSets struct {
	Amf3Gpp                       *Amf3GppAccessRegistration      `json:"amf3Gpp,omitempty"`
	AmfNon3Gpp                    *AmfNon3GppAccessRegistration   `json:"amfNon3Gpp,omitempty"`
	EeSubscriptions               []EeSubscription                `json:"eeSubscriptions,omitempty"`
	SmsfNon3GppAccess             *SmsfRegistration               `json:"smsfNon3GppAccess,omitempty"`
	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
	SmfRegistrations              []SmfRegistration               `json:"smfRegistrations,omitempty"`
	RoamingInfo                   *RoamingInfoUpdate              `json:"roamingInfo,omitempty"`
	SdmSubscriptions              []SdmSubscription               `json:"sdmSubscriptions,omitempty"`
	Smsf3GppAccess                *SmsfRegistration               `json:"smsf3GppAccess,omitempty"`
	IpSmGw                        *IpSmGwRegistration             `json:"ipSmGw,omitempty"`
	PeiInfo                       *PeiUpdateInfo                  `json:"peiInfo,omitempty"`
}
