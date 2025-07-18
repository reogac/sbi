/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	Dnn                             string                          `json:"dnn"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
}
