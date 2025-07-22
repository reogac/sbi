/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	PduSessionId            int                       `json:"pduSessionId"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	Dnn                     string                    `json:"dnn"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	Supi                    string                    `json:"supi"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
}
