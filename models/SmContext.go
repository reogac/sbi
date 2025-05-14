/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	Dnn                             string                          `json:"dnn"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
}
