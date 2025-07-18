/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GroupIdentifiers struct {
	UeIdList     []UeId   `json:"ueIdList,omitempty"`
	AllowedAfIds []string `json:"allowedAfIds,omitempty"`
	ExtGroupId   string   `json:"extGroupId,omitempty"`
	IntGroupId   string   `json:"intGroupId,omitempty"`
}
