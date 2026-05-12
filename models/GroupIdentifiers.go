/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GroupIdentifiers struct {
	ExtGroupId   string   `json:"extGroupId,omitempty"`
	IntGroupId   string   `json:"intGroupId,omitempty"`
	UeIdList     []UeId   `json:"ueIdList,omitempty"`
	AllowedAfIds []string `json:"allowedAfIds,omitempty"`
}
