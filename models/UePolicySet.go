/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	Pei                  string                                  `json:"pei,omitempty"`
	OsIds                []string                                `json:"osIds,omitempty"`
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
	SubscCats            []string                                `json:"subscCats,omitempty"`
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
}
