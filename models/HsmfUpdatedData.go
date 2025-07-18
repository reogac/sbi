/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdatedData struct {
	Ipv6MultiHomingInd              *bool                         `json:"ipv6MultiHomingInd,omitempty"`
	SessionAmbr                     *Ambr                         `json:"sessionAmbr,omitempty"`
	InterPlmnApiRoot                string                        `json:"interPlmnApiRoot,omitempty"`
	N4InfoExt2                      *N4Information                `json:"n4InfoExt2,omitempty"`
	SupportedFeatures               string                        `json:"supportedFeatures,omitempty"`
	HomeProvidedChargingId          string                        `json:"homeProvidedChargingId,omitempty"`
	Pti                             *int                          `json:"pti,omitempty"`
	IntraPlmnApiRoot                string                        `json:"intraPlmnApiRoot,omitempty"`
	N4InfoExt1                      *N4Information                `json:"n4InfoExt1,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem            `json:"qosFlowsSetupList,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	N4Info                          *N4Information                `json:"n4Info,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	DnaiList                        []string                      `json:"dnaiList,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	UpSecurity                      *UpSecurity                   `json:"upSecurity,omitempty"`
}
