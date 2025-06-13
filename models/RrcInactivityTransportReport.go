/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:12 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RrcInactivityTransportReport struct {
	RrcState int16         `json:"rrcState"`
	Loc      *UserLocation `json:"loc,omitempty"`
}
