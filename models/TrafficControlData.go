/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficControlData struct {
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	TcId                   string                 `json:"tcId"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
}
