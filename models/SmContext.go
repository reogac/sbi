/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
}
