/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	Dnn                     string                    `json:"dnn"`
	Pei                     string                    `json:"pei,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	Supi                    string                    `json:"supi"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
}
