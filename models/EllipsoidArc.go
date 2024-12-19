/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EllipsoidArc struct {
	IncludedAngle     int                     `json:"includedAngle"`
	Confidence        int                     `json:"confidence"`
	Point             GeographicalCoordinates `json:"point"`
	InnerRadius       int32                   `json:"innerRadius"`
	Shape             SupportedGADShapes      `json:"shape"`
	UncertaintyRadius float64                 `json:"uncertaintyRadius"`
	OffsetAngle       int                     `json:"offsetAngle"`
}
