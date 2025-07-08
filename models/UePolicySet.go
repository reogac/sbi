/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	Pei                  string                                  `json:"pei,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	SubscCats            []string                                `json:"subscCats,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	OsIds                []string                                `json:"osIds,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
}
