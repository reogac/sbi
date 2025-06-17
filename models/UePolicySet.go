/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	OsIds                []string                                `json:"osIds,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	SubscCats            []string                                `json:"subscCats,omitempty"`
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
	Pei                  string                                  `json:"pei,omitempty"`
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
}
