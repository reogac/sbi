/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	ContVer         *int                               `json:"contVer,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	Precedence      *int                               `json:"precedence,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
}
