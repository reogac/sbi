/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	RefChgData      []string                           `json:"refChgData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
}
