/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointUncertaintyCircle struct {
	Shape       SupportedGADShapes      `json:"shape"`
	Uncertainty float64                 `json:"uncertainty"`
	Point       GeographicalCoordinates `json:"point"`
}
