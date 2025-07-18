/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EllipsoidArc struct {
	UncertaintyRadius float64                 `json:"uncertaintyRadius"`
	OffsetAngle       int                     `json:"offsetAngle"`
	Shape             SupportedGADShapes      `json:"shape"`
	IncludedAngle     int                     `json:"includedAngle"`
	Confidence        int                     `json:"confidence"`
	Point             GeographicalCoordinates `json:"point"`
	InnerRadius       int32                   `json:"innerRadius"`
}
