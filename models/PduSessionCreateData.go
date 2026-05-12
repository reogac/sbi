/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	Dnn                             string                          `json:"dnn"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
}
