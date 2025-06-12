/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySet struct {
	UePolicySections     map[string]UePolicySection              `json:"uePolicySections,omitempty"`
	Upsis                []string                                `json:"upsis,omitempty"`
	AllowedRouteSelDescs map[string]PlmnRouteSelectionDescriptor `json:"allowedRouteSelDescs,omitempty"`
	AndspInd             *bool                                   `json:"andspInd,omitempty"`
	Pei                  string                                  `json:"pei,omitempty"`
	SuppFeat             string                                  `json:"suppFeat,omitempty"`
	PraInfos             map[string]PresenceInfo                 `json:"praInfos,omitempty"`
	OsIds                []string                                `json:"osIds,omitempty"`
	ResetIds             []string                                `json:"resetIds,omitempty"`
	SubscCats            []string                                `json:"subscCats,omitempty"`
}
