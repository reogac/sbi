/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RouteToLocation struct {
	Dnai        string            `json:"dnai"`
	RouteInfo   *RouteInformation `json:"routeInfo,omitempty"`
	RouteProfId string            `json:"routeProfId,omitempty"`
}
