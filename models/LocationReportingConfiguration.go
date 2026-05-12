/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationReportingConfiguration struct {
	OneTime         *bool            `json:"oneTime,omitempty"`
	Accuracy        LocationAccuracy `json:"accuracy,omitempty"`
	N3gppAccuracy   LocationAccuracy `json:"n3gppAccuracy,omitempty"`
	CurrentLocation bool             `json:"currentLocation"`
}
