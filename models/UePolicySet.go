/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	SubscCats            []string                                `json:"subscCats,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
	Pei                  string                                  `json:"pei,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	OsIds                []string                                `json:"osIds,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
}
