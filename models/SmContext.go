/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	Dnn                             string                          `json:"dnn"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
}
