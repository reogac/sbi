/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
}
