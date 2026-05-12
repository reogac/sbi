/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	RefChgData      []string                           `json:"refChgData,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
}
