/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationReportingConfiguration struct {
	CurrentLocation bool             `json:"currentLocation"`
	OneTime         *bool            `json:"oneTime,omitempty"`
	Accuracy        LocationAccuracy `json:"accuracy,omitempty"`
	N3gppAccuracy   LocationAccuracy `json:"n3gppAccuracy,omitempty"`
}
