/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:27 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RrcInactivityTransportReport struct {
	Loc      *UserLocation `json:"loc,omitempty"`
	RrcState int16         `json:"rrcState"`
}
