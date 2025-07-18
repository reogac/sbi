/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VsmfUpdateData struct {
	SmallDataRateControlEnabled *bool                         `json:"smallDataRateControlEnabled,omitempty"`
	QosFlowsAddModRequestList   []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	N4Info                      *N4Information                `json:"n4Info,omitempty"`
	N4InfoExt2                  *N4Information                `json:"n4InfoExt2,omitempty"`
	MaAcceptedInd               *bool                         `json:"maAcceptedInd,omitempty"`
	N4InfoExt1                  *N4Information                `json:"n4InfoExt1,omitempty"`
	Cause                       Cause                         `json:"cause,omitempty"`
	RequestIndication           RequestIndication             `json:"requestIndication"`
	QosFlowsRelRequestList      []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	QosMonitoringInfo           *QosMonitoringInfo            `json:"qosMonitoringInfo,omitempty"`
	SessionAmbr                 *Ambr                         `json:"sessionAmbr,omitempty"`
	EpsBearerInfo               []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	DnaiList                    []string                      `json:"dnaiList,omitempty"`
	N4InfoExt3                  *N4Information                `json:"n4InfoExt3,omitempty"`
	EpsPdnCnxInfo               *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	N9InactivityTimer           *int                          `json:"n9InactivityTimer,omitempty"`
	AssignEbiList               []Arp                         `json:"assignEbiList,omitempty"`
	ModifiedEbiList             []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	MaReleaseInd                MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	HsmfPduSessionUri           string                        `json:"hsmfPduSessionUri,omitempty"`
	NewSmfId                    string                        `json:"newSmfId,omitempty"`
	BackOffTimer                *int                          `json:"backOffTimer,omitempty"`
	AdditionalCnTunnelInfo      *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	N9DataForwardingInd         *bool                         `json:"n9DataForwardingInd,omitempty"`
	SupportedFeatures           string                        `json:"supportedFeatures,omitempty"`
	N1smCause                   string                        `json:"n1smCause,omitempty"`
	N1SmInfoToUe                *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	AlwaysOnGranted             *bool                         `json:"alwaysOnGranted,omitempty"`
	NewSmfPduSessionUri         string                        `json:"newSmfPduSessionUri,omitempty"`
	RevokeEbiList               []int                         `json:"revokeEbiList,omitempty"`
	Pti                         *int                          `json:"pti,omitempty"`
}
