/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MdtConfiguration struct {
	ReportIntervalNr         ReportIntervalNrMdt       `json:"reportIntervalNr,omitempty"`
	JobType                  JobType                   `json:"jobType"`
	ReportType               ReportTypeMdt             `json:"reportType,omitempty"`
	LoggingIntervalNr        LoggingIntervalNrMdt      `json:"loggingIntervalNr,omitempty"`
	CollectionPeriodRmmLte   CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	MeasurementPeriodLte     MeasurementPeriodLteMdt   `json:"measurementPeriodLte,omitempty"`
	AreaScope                *AreaScope                `json:"areaScope,omitempty"`
	MeasurementLteList       []string                  `json:"measurementLteList,omitempty"`
	ReportingTriggerList     []string                  `json:"reportingTriggerList,omitempty"`
	ReportInterval           ReportIntervalMdt         `json:"reportInterval,omitempty"`
	ReportAmount             ReportAmountMdt           `json:"reportAmount,omitempty"`
	EventThresholdRsrqNr     *int                      `json:"eventThresholdRsrqNr,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId                  `json:"mdtAllowedPlmnIdList,omitempty"`
	SensorMeasurementList    []string                  `json:"sensorMeasurementList,omitempty"`
	LoggingInterval          LoggingIntervalMdt        `json:"loggingInterval,omitempty"`
	PositioningMethod        PositioningMethodMdt      `json:"positioningMethod,omitempty"`
	EventThresholdRsrp       *int                      `json:"eventThresholdRsrp,omitempty"`
	InterFreqTargetList      []InterFreqTargetInfo     `json:"interFreqTargetList,omitempty"`
	MbsfnAreaList            []MbsfnArea               `json:"mbsfnAreaList,omitempty"`
	MeasurementNrList        []string                  `json:"measurementNrList,omitempty"`
	EventThresholdRsrpNr     *int                      `json:"eventThresholdRsrpNr,omitempty"`
	AddPositioningMethodList []string                  `json:"addPositioningMethodList,omitempty"`
	EventThresholdRsrq       *int                      `json:"eventThresholdRsrq,omitempty"`
	LoggingDuration          LoggingDurationMdt        `json:"loggingDuration,omitempty"`
	EventList                []string                  `json:"eventList,omitempty"`
	LoggingDurationNr        LoggingDurationNrMdt      `json:"loggingDurationNr,omitempty"`
	CollectionPeriodRmmNr    CollectionPeriodRmmNrMdt  `json:"collectionPeriodRmmNr,omitempty"`
}
