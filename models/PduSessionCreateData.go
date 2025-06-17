/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	AnType                          AccessType                      `json:"anType"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	Dnn                             string                          `json:"dnn"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
}
