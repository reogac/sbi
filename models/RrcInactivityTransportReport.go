/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RrcInactivityTransportReport struct {
	Loc      *UserLocation `json:"loc,omitempty"`
	RrcState int16         `json:"rrcState"`
}
