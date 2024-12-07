/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GroupIdentifiers struct {
	UeIdList     []UeId   `json:"ueIdList,omitempty"`
	AllowedAfIds []string `json:"allowedAfIds,omitempty"`
	ExtGroupId   string   `json:"extGroupId,omitempty"`
	IntGroupId   string   `json:"intGroupId,omitempty"`
}
