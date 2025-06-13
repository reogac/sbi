/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	Dnn                             string                          `json:"dnn"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
}
