/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SteeringMode struct {
	Active         AccessType         `json:"active,omitempty"`
	Standby        AccessType         `json:"standby,omitempty"`
	ThreeGLoad     *int               `json:"3gLoad,omitempty"`
	PrioAcc        AccessType         `json:"prioAcc,omitempty"`
	ThresValue     *ThresholdValue    `json:"thresValue,omitempty"`
	SteerModeInd   SteerModeIndicator `json:"steerModeInd,omitempty"`
	SteerModeValue SteerModeValue     `json:"steerModeValue"`
}
