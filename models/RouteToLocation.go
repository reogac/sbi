/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RouteToLocation struct {
	Dnai        string            `json:"dnai"`
	RouteInfo   *RouteInformation `json:"routeInfo,omitempty"`
	RouteProfId string            `json:"routeProfId,omitempty"`
}
