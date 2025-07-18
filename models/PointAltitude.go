/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointAltitude struct {
	Shape    SupportedGADShapes      `json:"shape"`
	Point    GeographicalCoordinates `json:"point"`
	Altitude float64                 `json:"altitude"`
}
