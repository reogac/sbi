/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GroupIdentifiers struct {
	UeIdList     []UeId   `json:"ueIdList,omitempty"`
	AllowedAfIds []string `json:"allowedAfIds,omitempty"`
	ExtGroupId   string   `json:"extGroupId,omitempty"`
	IntGroupId   string   `json:"intGroupId,omitempty"`
}
