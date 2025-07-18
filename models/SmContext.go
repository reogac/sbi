/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	Dnn                             string                          `json:"dnn"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
}
