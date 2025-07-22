/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	Dnn                             string                          `json:"dnn"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
}
