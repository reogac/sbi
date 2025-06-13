/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficControlData struct {
	TcId                   string                 `json:"tcId"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
}
