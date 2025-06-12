/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PlmnRestriction struct {
	CoreNetworkTypeRestrictions []string                `json:"coreNetworkTypeRestrictions,omitempty"`
	PrimaryRatRestrictions      []string                `json:"primaryRatRestrictions,omitempty"`
	SecondaryRatRestrictions    []string                `json:"secondaryRatRestrictions,omitempty"`
	RatRestrictions             []string                `json:"ratRestrictions,omitempty"`
	ForbiddenAreas              []Area                  `json:"forbiddenAreas,omitempty"`
	ServiceAreaRestriction      *ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`
}
