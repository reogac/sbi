/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficControlData struct {
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	TcId                   string                 `json:"tcId"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
}
