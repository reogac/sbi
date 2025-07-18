/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeographicArea struct {
	PointUncertaintyEllipse  *PointUncertaintyEllipse  `json:"PointUncertaintyEllipse,omitempty"`
	Polygon                  *Polygon                  `json:"Polygon,omitempty"`
	PointAltitude            *PointAltitude            `json:"PointAltitude,omitempty"`
	PointAltitudeUncertainty *PointAltitudeUncertainty `json:"PointAltitudeUncertainty,omitempty"`
	EllipsoidArc             *EllipsoidArc             `json:"EllipsoidArc,omitempty"`
	Point                    *Point                    `json:"Point,omitempty"`
	PointUncertaintyCircle   *PointUncertaintyCircle   `json:"PointUncertaintyCircle,omitempty"`
}
