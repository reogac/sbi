/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	Dnn                     string                    `json:"dnn"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	Supi                    string                    `json:"supi"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
}
