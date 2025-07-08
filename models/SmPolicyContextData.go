/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	Online                  *bool                     `json:"online,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	Supi                    string                    `json:"supi"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	Dnn                     string                    `json:"dnn"`
	NotificationUri         string                    `json:"notificationUri"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
}
