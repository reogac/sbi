/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	OsIds                []string                                `json:"osIds,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	Pei                  string                                  `json:"pei,omitempty"`
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
	SubscCats            []string                                `json:"subscCats,omitempty"`
}
