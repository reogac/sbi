/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationReportingConfiguration struct {
	Accuracy        LocationAccuracy `json:"accuracy,omitempty"`
	N3gppAccuracy   LocationAccuracy `json:"n3gppAccuracy,omitempty"`
	CurrentLocation bool             `json:"currentLocation"`
	OneTime         *bool            `json:"oneTime,omitempty"`
}
