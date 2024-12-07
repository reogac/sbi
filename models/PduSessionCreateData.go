/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	Dnn                             string                          `json:"dnn"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
}
