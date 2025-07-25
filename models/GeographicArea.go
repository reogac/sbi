/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeographicArea struct {
	Point                    *Point                    `json:"Point,omitempty"`
	PointUncertaintyCircle   *PointUncertaintyCircle   `json:"PointUncertaintyCircle,omitempty"`
	PointUncertaintyEllipse  *PointUncertaintyEllipse  `json:"PointUncertaintyEllipse,omitempty"`
	Polygon                  *Polygon                  `json:"Polygon,omitempty"`
	PointAltitude            *PointAltitude            `json:"PointAltitude,omitempty"`
	PointAltitudeUncertainty *PointAltitudeUncertainty `json:"PointAltitudeUncertainty,omitempty"`
	EllipsoidArc             *EllipsoidArc             `json:"EllipsoidArc,omitempty"`
}
