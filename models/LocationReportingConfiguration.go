/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationReportingConfiguration struct {
	N3gppAccuracy   LocationAccuracy `json:"n3gppAccuracy,omitempty"`
	CurrentLocation bool             `json:"currentLocation"`
	OneTime         *bool            `json:"oneTime,omitempty"`
	Accuracy        LocationAccuracy `json:"accuracy,omitempty"`
}
