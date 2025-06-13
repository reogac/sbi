/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeCampingRep struct {
	AccessType          AccessType                `json:"accessType,omitempty"`
	RatType             RatType                   `json:"ratType,omitempty"`
	ServNfId            *ServingNfIdentity        `json:"servNfId,omitempty"`
	ServingNetwork      *PlmnIdNid                `json:"servingNetwork,omitempty"`
	UserLocationInfo    *UserLocation             `json:"userLocationInfo,omitempty"`
	UeTimeZone          string                    `json:"ueTimeZone,omitempty"`
	NetLocAccSupp       NetLocAccessSupport       `json:"netLocAccSupp,omitempty"`
	SatBackhaulCategory SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
}
