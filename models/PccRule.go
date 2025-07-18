/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	RefTcData       []string                           `json:"refTcData,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
}
