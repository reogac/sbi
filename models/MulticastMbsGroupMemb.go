/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MulticastMbsGroupMemb struct {
	AfInstanceId            string   `json:"afInstanceId,omitempty"`
	InternalGroupIdentifier string   `json:"internalGroupIdentifier,omitempty"`
	MulticastGroupMemb      []string `json:"multicastGroupMemb"`
}
