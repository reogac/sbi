/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeCampingRep struct {
	RatType             RatType                   `json:"ratType,omitempty"`
	ServNfId            *ServingNfIdentity        `json:"servNfId,omitempty"`
	ServingNetwork      *PlmnIdNid                `json:"servingNetwork,omitempty"`
	UserLocationInfo    *UserLocation             `json:"userLocationInfo,omitempty"`
	UeTimeZone          string                    `json:"ueTimeZone,omitempty"`
	NetLocAccSupp       NetLocAccessSupport       `json:"netLocAccSupp,omitempty"`
	SatBackhaulCategory SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	AccessType          AccessType                `json:"accessType,omitempty"`
}
