/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeCampingRep struct {
	UserLocationInfo    *UserLocation             `json:"userLocationInfo,omitempty"`
	UeTimeZone          string                    `json:"ueTimeZone,omitempty"`
	NetLocAccSupp       NetLocAccessSupport       `json:"netLocAccSupp,omitempty"`
	SatBackhaulCategory SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	AccessType          AccessType                `json:"accessType,omitempty"`
	RatType             RatType                   `json:"ratType,omitempty"`
	ServNfId            *ServingNfIdentity        `json:"servNfId,omitempty"`
	ServingNetwork      *PlmnIdNid                `json:"servingNetwork,omitempty"`
}
