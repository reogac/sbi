/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TscaiInputContainer struct {
	Periodicity      *int   `json:"periodicity,omitempty"`
	BurstArrivalTime string `json:"burstArrivalTime,omitempty"`
	SurTimeInNumMsg  *int   `json:"surTimeInNumMsg,omitempty"`
	SurTimeInTime    *int   `json:"surTimeInTime,omitempty"`
}
