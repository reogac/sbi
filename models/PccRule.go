/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	Precedence      *int                               `json:"precedence,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
}
