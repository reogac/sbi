/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RrcInactivityTransportReport struct {
	Loc      *UserLocation `json:"loc,omitempty"`
	RrcState int16         `json:"rrcState"`
}
