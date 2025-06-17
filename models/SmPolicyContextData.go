/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:58 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	AccessType              AccessType                `json:"accessType,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	Supi                    string                    `json:"supi"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	Dnn                     string                    `json:"dnn"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
}
