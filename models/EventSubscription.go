/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
}
