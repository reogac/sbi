/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	Precedence      *int                               `json:"precedence,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
}
