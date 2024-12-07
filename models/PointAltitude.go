/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointAltitude struct {
	Shape    SupportedGADShapes      `json:"shape"`
	Altitude float64                 `json:"altitude"`
	Point    GeographicalCoordinates `json:"point"`
}
