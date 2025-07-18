/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
}
