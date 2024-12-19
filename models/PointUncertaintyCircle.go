/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointUncertaintyCircle struct {
	Point       GeographicalCoordinates `json:"point"`
	Uncertainty float64                 `json:"uncertainty"`
	Shape       SupportedGADShapes      `json:"shape"`
}
