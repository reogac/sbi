/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	Supi                    string                    `json:"supi"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	Dnn                     string                    `json:"dnn"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
}
