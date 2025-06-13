/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowItem struct {
	CurrentQosProfileIndex *int       `json:"currentQosProfileIndex,omitempty"`
	NullQoSProfileIndex    *bool      `json:"nullQoSProfileIndex,omitempty"`
	NgApCause              *NgApCause `json:"ngApCause,omitempty"`
	Qfi                    int        `json:"qfi"`
	Cause                  Cause      `json:"cause,omitempty"`
}
