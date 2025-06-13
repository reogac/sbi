/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
}
