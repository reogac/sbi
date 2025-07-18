/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreatedData struct {
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	UeIpv6InterfaceId               string                          `json:"ueIpv6InterfaceId,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	SessionAmbr                     *Ambr                           `json:"sessionAmbr,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData                `json:"n1SmInfoToUe,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	HomeProvidedChargingId          string                          `json:"homeProvidedChargingId,omitempty"`
	AdditionalSnssai                *Snssai                         `json:"additionalSnssai,omitempty"`
	CnTunnelInfo                    *TunnelInfo                     `json:"cnTunnelInfo,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem              `json:"qosFlowsSetupList,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	SmallDataRateControlEnabled     *bool                           `json:"smallDataRateControlEnabled,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	Ipv6MultiHomingInd              *bool                           `json:"ipv6MultiHomingInd,omitempty"`
	HcnTunnelInfo                   *TunnelInfo                     `json:"hcnTunnelInfo,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	SscMode                         string                          `json:"sscMode"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	MaAcceptedInd                   *bool                           `json:"maAcceptedInd,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
}
