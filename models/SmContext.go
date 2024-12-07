/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	Dnn                             string                          `json:"dnn"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
}
