/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PlmnRestriction struct {
	RatRestrictions             []string                `json:"ratRestrictions,omitempty"`
	ForbiddenAreas              []Area                  `json:"forbiddenAreas,omitempty"`
	ServiceAreaRestriction      *ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions []string                `json:"coreNetworkTypeRestrictions,omitempty"`
	PrimaryRatRestrictions      []string                `json:"primaryRatRestrictions,omitempty"`
	SecondaryRatRestrictions    []string                `json:"secondaryRatRestrictions,omitempty"`
}
