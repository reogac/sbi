/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TscaiInputContainer struct {
	Periodicity      *int   `json:"periodicity,omitempty"`
	BurstArrivalTime string `json:"burstArrivalTime,omitempty"`
	SurTimeInNumMsg  *int   `json:"surTimeInNumMsg,omitempty"`
	SurTimeInTime    *int   `json:"surTimeInTime,omitempty"`
}
