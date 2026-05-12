/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficControlData struct {
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	TcId                   string                 `json:"tcId"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
}
