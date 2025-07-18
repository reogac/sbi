/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeCampingRep struct {
	ServingNetwork      *PlmnIdNid                `json:"servingNetwork,omitempty"`
	UserLocationInfo    *UserLocation             `json:"userLocationInfo,omitempty"`
	UeTimeZone          string                    `json:"ueTimeZone,omitempty"`
	NetLocAccSupp       NetLocAccessSupport       `json:"netLocAccSupp,omitempty"`
	SatBackhaulCategory SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	AccessType          AccessType                `json:"accessType,omitempty"`
	RatType             RatType                   `json:"ratType,omitempty"`
	ServNfId            *ServingNfIdentity        `json:"servNfId,omitempty"`
}
