/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosRequirement struct {
	FiveQi  *int            `json:"5qi,omitempty"`
	GfbrUl  string          `json:"gfbrUl,omitempty"`
	GfbrDl  string          `json:"gfbrDl,omitempty"`
	ResType QosResourceType `json:"resType,omitempty"`
	Pdb     *int            `json:"pdb,omitempty"`
	Per     string          `json:"per,omitempty"`
}
