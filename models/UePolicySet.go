/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	SubscCats            []string                                `json:"subscCats,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	Pei                  string                                  `json:"pei,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	OsIds                []string                                `json:"osIds,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
}
