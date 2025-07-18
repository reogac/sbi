/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	Dnn                             string                          `json:"dnn"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
}
