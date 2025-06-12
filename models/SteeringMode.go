/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SteeringMode struct {
	SteerModeValue SteerModeValue     `json:"steerModeValue"`
	Active         AccessType         `json:"active,omitempty"`
	Standby        AccessType         `json:"standby,omitempty"`
	ThreeGLoad     *int               `json:"3gLoad,omitempty"`
	PrioAcc        AccessType         `json:"prioAcc,omitempty"`
	ThresValue     *ThresholdValue    `json:"thresValue,omitempty"`
	SteerModeInd   SteerModeIndicator `json:"steerModeInd,omitempty"`
}
