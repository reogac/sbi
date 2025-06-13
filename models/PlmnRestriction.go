/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PlmnRestriction struct {
	ServiceAreaRestriction      *ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions []string                `json:"coreNetworkTypeRestrictions,omitempty"`
	PrimaryRatRestrictions      []string                `json:"primaryRatRestrictions,omitempty"`
	SecondaryRatRestrictions    []string                `json:"secondaryRatRestrictions,omitempty"`
	RatRestrictions             []string                `json:"ratRestrictions,omitempty"`
	ForbiddenAreas              []Area                  `json:"forbiddenAreas,omitempty"`
}
