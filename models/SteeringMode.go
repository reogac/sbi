/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SteeringMode struct {
	ThresValue     *ThresholdValue    `json:"thresValue,omitempty"`
	SteerModeInd   SteerModeIndicator `json:"steerModeInd,omitempty"`
	SteerModeValue SteerModeValue     `json:"steerModeValue"`
	Active         AccessType         `json:"active,omitempty"`
	Standby        AccessType         `json:"standby,omitempty"`
	ThreeGLoad     *int               `json:"3gLoad,omitempty"`
	PrioAcc        AccessType         `json:"prioAcc,omitempty"`
}
