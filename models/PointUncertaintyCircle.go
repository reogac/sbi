/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointUncertaintyCircle struct {
	Shape       SupportedGADShapes      `json:"shape"`
	Point       GeographicalCoordinates `json:"point"`
	Uncertainty float64                 `json:"uncertainty"`
}
