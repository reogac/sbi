/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointUncertaintyEllipse struct {
	Confidence         int                     `json:"confidence"`
	Point              GeographicalCoordinates `json:"point"`
	Shape              SupportedGADShapes      `json:"shape"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
}
