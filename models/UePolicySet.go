/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	Pei                  string                                  `json:"pei,omitempty"`
	OsIds                []string                                `json:"osIds,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
	SubscCats            []string                                `json:"subscCats,omitempty"`
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
}
