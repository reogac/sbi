/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnQuerySmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmsfContext3gppParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuerySmsfContext3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetIndividualSharedData(ctx sbi.RequestContext, prod Producer) {
	var params GetIndividualSharedDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'sharedDataId'
	params.SharedDataId = ctx.Param("sharedDataId")
	if len(params.SharedDataId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sharedDataId is required"), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetIndividualSharedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnModifysubscriptionDataSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifysubscriptionDataSubscriptionParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysubscriptionDataSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryLcsMoData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryLcsMoDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryLcsMoData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnModifyOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyOperSpecData(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifySmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySmfSubscriptionInfoParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyHssSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyHssSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroup(gpsis)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryCoverageRestrictionData(ctx sbi.RequestContext, prod Producer) {
	var params QueryCoverageRestrictionDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp := prod.HandleQueryCoverageRestrictionData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnGetAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params GetAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyAuthenticationSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAuthenticationSubscriptionParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAuthenticationSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetNiddAuthorizationInfo(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyNiddAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyNiddAuthorizationInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryUeSubscribedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryUeSubscribedDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)), nil)
			return
		}
	}

	// read 'serving-plmn'
	params.ServingPlmn = ctx.Param("serving-plmn")

	// call application handler
	rsp := prod.HandleQueryUeSubscribedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryeeSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QueryeeSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQueryeeSubscription(&params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQuerySmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmsfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmsfContextNon3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateSdmSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SdmSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateSdmSubscriptions(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnQueryV2xData(ctx sbi.RequestContext, prod Producer) {
	var params QueryV2xDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp := prod.HandleQueryV2xData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuery5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GmbsGroup(gpsis)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnSubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SubscriptionDataSubscriptions)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleSubscriptionDataSubscriptions(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnQueryNssaiAck(ctx sbi.RequestContext, prod Producer) {
	var params QueryNssaiAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryNssaiAck(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifysdmSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysdmSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveHssSDMSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveHssSDMSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveHssSDMSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryContextData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryContextDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'context-dataset-names'
	contextDatasetNamesStr := ctx.Param("context-dataset-names")
	if len(contextDatasetNamesStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "context-dataset-names is required"), nil)
		return
	}

	if params.ContextDatasetNames, err = models.ArrayOfStringFromString(contextDatasetNamesStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse context-dataset-names failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryContextData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryLcsPrivacyData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryLcsPrivacyDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp := prod.HandleQueryLcsPrivacyData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryAuthUPU(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthUPUParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthUPU(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryPPData(ctx sbi.RequestContext, prod Producer) {
	var params QueryPPDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params QueryEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryEeGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemoveAmfSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveAmfSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveAmfSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateHSSSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateHSSSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HssSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateHSSSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveServiceSpecificAuthorizationInfoParams

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveServiceSpecificAuthorizationInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryRoamingInformation(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryRoamingInformation(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuery5GMbsGroupInternal(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"), nil)
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupInternal(internalGroupIds)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnUpdateRoamingInformation(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.RoamingInfoUpdate)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleUpdateRoamingInformation(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryCagAck(ctx sbi.RequestContext, prod Producer) {
	var params QueryCagAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryCagAck(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryAmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryAmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmfContextNon3gppParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryAmfContextNon3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmsData(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmsDataParams

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQuerySmsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnGetAmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetAmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdatesdmsubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdatesdmsubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SdmSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdatesdmsubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemovesdmSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemovesdmSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleRemovesdmSubscriptions(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveAmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveAmfGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteAuthenticationStatus(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteSmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContext3gpp(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetGroupIdentifiers(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetGroupIdentifiersParams

	// read 'ext-group-id'
	params.ExtGroupId = ctx.Param("ext-group-id")

	// read 'int-group-id'
	params.IntGroupId = ctx.Param("int-group-id")

	// read 'ue-id-ind'
	ueIdIndStr := ctx.Param("ue-id-ind")
	if len(ueIdIndStr) > 0 {
		var ueIdIndTmp bool
		if ueIdIndTmp, err = models.BoolFromString(ueIdIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ue-id-ind failed: %+v", err)), nil)
			return
		}
		params.UeIdInd = &ueIdIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetGroupIdentifiers(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreate5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.FiveGVnGroupConfiguration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GVnGroup(externalGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateNIDDAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NiddAuthorizationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateNIDDAuthorizationInfo(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGet5GVnGroupConfiguration(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGet5GVnGroupConfiguration(externalGroupId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmfSelectData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmfSelectDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQuerySmfSelectData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AmfNon3GppAccessRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateAmfContextNon3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryMessageWaitingDataParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryMessageWaitingData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetPPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var params GetPPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetPPDataEntry(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQueryeesubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryeesubscriptionsParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'event-types'
	eventTypesStr := ctx.Param("event-types")
	if len(eventTypesStr) > 0 {
		if params.EventTypes, err = models.ArrayOfStringFromString(eventTypesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse event-types failed: %+v", err)), nil)
			return
		}
	}

	// read 'nf-identifiers'
	nfIdentifiersStr := ctx.Param("nf-identifiers")
	if len(nfIdentifiersStr) > 0 {
		if params.NfIdentifiers, err = models.ArrayOfNfIdentifierFromString(nfIdentifiersStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse nf-identifiers failed: %+v", err)), nil)
			return
		}
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryeesubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateAMFSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateAMFSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetIdentityData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetIdentityDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'app-port-id'
	appPortIdStr := ctx.Param("app-port-id")
	if len(appPortIdStr) > 0 {
		if params.AppPortId, err = models.AppPortIdFromString(appPortIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse app-port-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetIdentityData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQueryAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmfContext3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAmfContext3gpp(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdateAuthenticationSoR(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateAuthenticationSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateAuthenticationSoR(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySmfRegList(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmfRegListParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmfRegList(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryUserConsentData(ctx sbi.RequestContext, prod Producer) {
	var params QueryUserConsentDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ucPurpose'
	params.UcPurpose = ctx.Param("ucPurpose")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryUserConsentData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateOrUpdateSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOrUpdateSmfRegistrationParams

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdateSmfRegistration(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteOperSpecData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteOperSpecData(ueId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteSmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContextNon3gpp(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyEesubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyEesubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEesubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyEeGroupSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyEeGroupSubscriptionParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEeGroupSubscription(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateHSSSDMSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateHSSSDMSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HssSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateHSSSDMSubscriptions(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveHssSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveHssSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveHssSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetHssSDMSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetHssSDMSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSDMSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemovesubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemovesubscriptionDataSubscriptions(subsId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateSmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetMulticastMbsGroupMemb(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetMulticastMbsGroupMemb(externalGroupId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationStatus(ueId, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryIndividualAuthenticationStatusParams

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryIndividualAuthenticationStatus(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params AmfContext3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContext3gpp(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryOperSpecData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuerySmsMngData(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmsMngDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp := prod.HandleQuerySmsMngData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryEeGroupSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QueryEeGroupSubscriptionParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQueryEeGroupSubscription(&params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnRemoveMultipleSubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params RemoveMultipleSubscriptionDataSubscriptionsParams

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"), nil)
		return
	}

	// read 'nf-instance-id'
	params.NfInstanceId = ctx.Param("nf-instance-id")

	// read 'delete-all-nfs'
	deleteAllNfsStr := ctx.Param("delete-all-nfs")
	if len(deleteAllNfsStr) > 0 {
		var deleteAllNfsTmp bool
		if deleteAllNfsTmp, err = models.BoolFromString(deleteAllNfsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse delete-all-nfs failed: %+v", err)), nil)
			return
		}
		params.DeleteAllNfs = &deleteAllNfsTmp
	}

	// read 'implicit-unsubscribe-indication'
	implicitUnsubscribeIndicationStr := ctx.Param("implicit-unsubscribe-indication")
	if len(implicitUnsubscribeIndicationStr) > 0 {
		var implicitUnsubscribeIndicationTmp bool
		if implicitUnsubscribeIndicationTmp, err = models.BoolFromString(implicitUnsubscribeIndicationStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse implicit-unsubscribe-indication failed: %+v", err)), nil)
			return
		}
		params.ImplicitUnsubscribeIndication = &implicitUnsubscribeIndicationTmp
	}

	// call application handler
	prod.HandleRemoveMultipleSubscriptionDataSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5GVnGroupInternal(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"), nil)
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroupInternal(internalGroupIds)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateAmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDelete5GmbsGroup(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDelete5GmbsGroup(externalGroupId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveNiddAuthorizationInfo(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModify5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Modify5GmbsGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GmbsGroup(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params AmfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContextNon3gpp(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyAmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAmfSubscriptionInfoParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuerySubscriptionDataSubscriptions(subsId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.ServiceSpecificAuthorizationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateServiceSpecificAuthorizationInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyServiceSpecificAuthorizationInfoParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyServiceSpecificAuthorizationInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreate5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.MulticastMbsGroupMemb)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GmbsGroup(externalGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQueryAuthSoR(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthSoR(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateAuthenticationSoR(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAuthenticationSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SorData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationSoR(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateEeSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateEeSubscriptions(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnUpdateEesubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateEesubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateEesubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateEeGroupSubscriptions(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveEeGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetOdbData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetOdbData(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateOrUpdateNssaiAck(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOrUpdateNssaiAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NssaiAckData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateOrUpdateNssaiAck(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryIpSmGwContextParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryIpSmGwContext(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnDeleteMessageWaitingData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteMessageWaitingData(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetMultiplePPDataEntries(ctx sbi.RequestContext, prod Producer) {
	var params GetMultiplePPDataEntriesParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetMultiplePPDataEntries(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnRemoveSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveSmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveSmfGroupSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifySmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfGroupSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryTraceData(ctx sbi.RequestContext, prod Producer) {
	var params QueryTraceDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryTraceData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuerySmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmfRegistration(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateSmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmsfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContext3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateSmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmsfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContextNon3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteIpSmGwContext(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteIpSmGwContext(ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetppData(ctx sbi.RequestContext, prod Producer) {
	var params GetppDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetppData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQueryEEData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryEEDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryEEData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerysdmSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QuerysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQuerysdmSubscription(&params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQueryAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAuthenticationStatusParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryAuthenticationStatus(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateCagUpdateAck(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateCagUpdateAckParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.CagAckData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateCagUpdateAck(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryProvisionedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryProvisionedDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryProvisionedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.MessageWaitingData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateMessageWaitingData(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeletePPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var params DeletePPDataEntryParams

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeletePPDataEntry(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateIndividualAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateIndividualAuthenticationStatus(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateOperSpecData(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetHssSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetHssSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueGroupId'
	var ueGroupId string
	ueGroupId = ctx.Param("ueGroupId")
	if len(ueGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateEeGroupSubscriptions(ueGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnQueryUeLocation(ctx sbi.RequestContext, prod Producer) {
	var params QueryUeLocationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryUeLocation(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.IpSmGwRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateIpSmGwContext(ueId, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveSmfSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveSmfSubscriptionsInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveSmfSubscriptionsInfo(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerysdmsubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params QuerysdmsubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerysdmsubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyHssSDMSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyHssSDMSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSDMSubscriptionInfo(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSharedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSharedDataParams

	// read 'shared-data-ids'
	sharedDataIdsStr := ctx.Param("shared-data-ids")
	if len(sharedDataIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "shared-data-ids is required"), nil)
		return
	}

	if params.SharedDataIds, err = models.ArrayOfStringFromString(sharedDataIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse shared-data-ids failed: %+v", err)), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSharedData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnDelete5GVnGroup(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDelete5GVnGroup(externalGroupId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetServiceSpecificAuthorizationInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateAuthenticationUPU(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAuthenticationUPUParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UpuData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationUPU(&params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveeeSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveeeSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveeeSubscriptions(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSSAuData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSSAuDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"), nil)
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
		return
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"), nil)
		return
	}

	// read 'mtc-provider-information'
	params.MtcProviderInformation = ctx.Param("mtc-provider-information")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	prob := prod.HandleGetSSAuData(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQuery5GMbsGroupPPData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Query5GMbsGroupPPDataParams

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnDeleteIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var params DeleteIndividualAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteIndividualAuthenticationStatus(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetNiddAuData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetNiddAuDataParams

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"), nil)
		return
	}

	// read 'mtc-provider-information'
	params.MtcProviderInformation = ctx.Param("mtc-provider-information")
	if len(params.MtcProviderInformation) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "mtc-provider-information is required"), nil)
		return
	}

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"), nil)
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleGetNiddAuData(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQueryLcsBcaData(ctx sbi.RequestContext, prod Producer) {
	var params QueryLcsBcaDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryLcsBcaData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuery5mbsData(ctx sbi.RequestContext, prod Producer) {
	var params Query5mbsDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQuery5mbsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnUpdateSmfContext(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateSmfContextParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSmfContext(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModify5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Modify5GVnGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GVnGroup(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryPeiInformation(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryPeiInformation(ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateOrUpdatePeiInformation(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PeiUpdateInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdatePeiInformation(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params GetSmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfGroupSubscriptions(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryAuthSubsData(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthSubsDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryAuthSubsData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryPorseData(ctx sbi.RequestContext, prod Producer) {
	var params QueryPorseDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp := prod.HandleQueryPorseData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuerySubsToNotify(ctx sbi.RequestContext, prod Producer) {
	var params QuerySubsToNotifyParams

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySubsToNotify(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
			return
		}
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQuerySmData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnDeleteSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params DeleteSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmfRegistration(&params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleModifyIpSmGwContext(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleModifyMessageWaitingData(ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreatePPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreatePPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PpDataEntry)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreatePPDataEntry(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetSmfSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfSubscriptionInfo(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateSMFSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateSMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSMFSubscriptions(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.Amf3GppAccessRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateAmfContext3gpp(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyPpData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyPpDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyPpData(&params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryGroupEEData(ctx sbi.RequestContext, prod Producer) {
	var params QueryGroupEEDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryGroupEEData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuery5GVNGroupPPData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Query5GVNGroupPPDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuery5GVNGroupPPData(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

type Producer interface {
	HandleQuerySmsfContext3gpp(*QuerySmsfContext3gppParams) *models.SmsfRegistration

	HandleGetIndividualSharedData(*GetIndividualSharedDataParams) (map[string]string, *models.SharedData, *models.ProblemDetails)

	HandleModifysubscriptionDataSubscription(*ModifysubscriptionDataSubscriptionParams, *[]models.PatchItem) (*models.Schema, *models.ProblemDetails)

	HandleQueryLcsMoData(*QueryLcsMoDataParams) (map[string]string, *models.LcsMoData)

	HandleModifyOperSpecData(*ModifyOperSpecDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifySmfSubscriptionInfo(*ModifySmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifyHssSubscriptionInfo(*ModifyHssSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQuery5GVnGroup([]string) *map[string]models.FiveGVnGroupConfiguration

	HandleQueryCoverageRestrictionData(*QueryCoverageRestrictionDataParams) (map[string]string, *models.EnhancedCoverageRestrictionData)

	HandleGetAmfGroupSubscriptions(*GetAmfGroupSubscriptionsParams) *[]models.AmfSubscriptionInfo

	HandleModifyAuthenticationSubscription(*ModifyAuthenticationSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetNiddAuthorizationInfo(string) *models.NiddAuthorizationInfo

	HandleModifyNiddAuthorizationInfo(*ModifyNiddAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryUeSubscribedData(*QueryUeSubscribedDataParams) *models.UeSubscribedDataSets

	HandleModifyAmfGroupSubscriptions(*ModifyAmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryeeSubscription(*QueryeeSubscriptionParams)

	HandleQuerySmsfContextNon3gpp(*QuerySmsfContextNon3gppParams) *models.SmsfRegistration

	HandleCreateSdmSubscriptions(string, *models.SdmSubscription) (map[string]string, *models.SdmSubscription)

	HandleQueryV2xData(*QueryV2xDataParams) (map[string]string, *models.V2xSubscriptionData)

	HandleQuery5GmbsGroup([]string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleSubscriptionDataSubscriptions(*models.SubscriptionDataSubscriptions) (map[string]string, *models.SubscriptionDataSubscriptions)

	HandleQueryNssaiAck(*QueryNssaiAckParams) *models.NssaiAckData

	HandleModifysdmSubscription(*ModifysdmSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleRemoveHssSDMSubscriptionsInfo(*RemoveHssSDMSubscriptionsInfoParams)

	HandleQueryContextData(*QueryContextDataParams) *models.ContextDataSets

	HandleQueryLcsPrivacyData(*QueryLcsPrivacyDataParams) (map[string]string, *models.LcsPrivacyData)

	HandleQueryAuthUPU(*QueryAuthUPUParams) *models.UpuData

	HandleQueryPPData(*QueryPPDataParams) *models.PpProfileData

	HandleQueryEeGroupSubscriptions(*QueryEeGroupSubscriptionsParams) *[]models.EeSubscription

	HandleRemoveAmfSubscriptionsInfo(*RemoveAmfSubscriptionsInfoParams)

	HandleCreateHSSSubscriptions(*CreateHSSSubscriptionsParams, *models.HssSubscriptionInfo) *models.HssSubscriptionInfo

	HandleRemoveServiceSpecificAuthorizationInfo(*RemoveServiceSpecificAuthorizationInfoParams)

	HandleQueryRoamingInformation(string) *models.RoamingInfoUpdate

	HandleQuery5GMbsGroupInternal([]string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleUpdateRoamingInformation(string, *models.RoamingInfoUpdate) *models.RoamingInfoUpdate

	HandleQueryCagAck(*QueryCagAckParams) *models.CagAckData

	HandleQueryAmData(*QueryAmDataParams) (map[string]string, *models.AccessAndMobilitySubscriptionData)

	HandleQueryAmfContextNon3gpp(*QueryAmfContextNon3gppParams) *models.AmfNon3GppAccessRegistration

	HandleQuerySmsData(*QuerySmsDataParams) (map[string]string, *models.SmsSubscriptionData)

	HandleGetAmfSubscriptionInfo(*GetAmfSubscriptionInfoParams) *[]models.AmfSubscriptionInfo

	HandleUpdatesdmsubscriptions(*UpdatesdmsubscriptionsParams, *models.SdmSubscription) *models.ProblemDetails

	HandleRemovesdmSubscriptions(*RemovesdmSubscriptionsParams) *models.ProblemDetails

	HandleRemoveAmfGroupSubscriptions(*RemoveAmfGroupSubscriptionsParams)

	HandleDeleteAuthenticationStatus(string)

	HandleDeleteSmsfContext3gpp(string)

	HandleGetGroupIdentifiers(*GetGroupIdentifiersParams) (*models.GroupIdentifiers, *models.ProblemDetails)

	HandleCreate5GVnGroup(string, *models.FiveGVnGroupConfiguration) (*models.FiveGVnGroupConfiguration, *models.ProblemDetails)

	HandleCreateNIDDAuthorizationInfo(string, *models.NiddAuthorizationInfo) *models.NiddAuthorizationInfo

	HandleGet5GVnGroupConfiguration(string) *models.FiveGVnGroupConfiguration

	HandleQuerySmfSelectData(*QuerySmfSelectDataParams) (map[string]string, *models.SmfSelectionSubscriptionData)

	HandleCreateAmfContextNon3gpp(string, *models.AmfNon3GppAccessRegistration) (map[string]string, *models.Amf3GppAccessRegistration)

	HandleQueryMessageWaitingData(*QueryMessageWaitingDataParams) *models.MessageWaitingData

	HandleGetPPDataEntry(*GetPPDataEntryParams) (*models.PpDataEntry, *models.ProblemDetails)

	HandleQueryeesubscriptions(*QueryeesubscriptionsParams) *[]models.EeSubscriptionExt

	HandleCreateAMFSubscriptions(*CreateAMFSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleGetIdentityData(*GetIdentityDataParams) (map[string]string, *models.IdentityData, *models.ProblemDetails)

	HandleQueryAmfContext3gpp(*QueryAmfContext3gppParams) *models.Amf3GppAccessRegistration

	HandleUpdateAuthenticationSoR(*UpdateAuthenticationSoRParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQuerySmfRegList(*QuerySmfRegListParams) *[]models.SmfRegistration

	HandleQueryUserConsentData(*QueryUserConsentDataParams) (map[string]string, *models.UcSubscriptionData)

	HandleCreateOrUpdateSmfRegistration(*CreateOrUpdateSmfRegistrationParams, *models.SmfRegistration) *models.SmfRegistration

	HandleDeleteOperSpecData(string) *models.ProblemDetails

	HandleDeleteSmsfContextNon3gpp(string)

	HandleModifyEesubscription(*ModifyEesubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifyEeGroupSubscription(*ModifyEeGroupSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateHSSSDMSubscriptions(*CreateHSSSDMSubscriptionsParams, *models.HssSubscriptionInfo)

	HandleRemoveHssSubscriptionsInfo(*RemoveHssSubscriptionsInfoParams)

	HandleGetHssSDMSubscriptionInfo(*GetHssSDMSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleRemovesubscriptionDataSubscriptions(string)

	HandleCreateSmfGroupSubscriptions(*CreateSmfGroupSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleGetMulticastMbsGroupMemb(string) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleCreateAuthenticationStatus(string, *models.AuthEvent)

	HandleQueryIndividualAuthenticationStatus(*QueryIndividualAuthenticationStatusParams) *models.AuthEvent

	HandleAmfContext3gpp(*AmfContext3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryOperSpecData(*QueryOperSpecDataParams) (map[string]string, *map[string]models.OperatorSpecificDataContainer)

	HandleQuerySmsMngData(*QuerySmsMngDataParams) (map[string]string, *models.SmsManagementSubscriptionData)

	HandleQueryEeGroupSubscription(*QueryEeGroupSubscriptionParams)

	HandleRemoveMultipleSubscriptionDataSubscriptions(*RemoveMultipleSubscriptionDataSubscriptionsParams)

	HandleQuery5GVnGroupInternal([]string) *map[string]models.FiveGVnGroupConfiguration

	HandleCreateAmfGroupSubscriptions(*CreateAmfGroupSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleDelete5GmbsGroup(string) *models.ProblemDetails

	HandleRemoveNiddAuthorizationInfo(string)

	HandleModify5GmbsGroup(*Modify5GmbsGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleAmfContextNon3gpp(*AmfContextNon3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifyAmfSubscriptionInfo(*ModifyAmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQuerySubscriptionDataSubscriptions(string) *models.SubscriptionDataSubscriptions

	HandleCreateServiceSpecificAuthorizationInfo(*CreateServiceSpecificAuthorizationInfoParams, *models.ServiceSpecificAuthorizationInfo) *models.ServiceSpecificAuthorizationInfo

	HandleModifyServiceSpecificAuthorizationInfo(*ModifyServiceSpecificAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreate5GmbsGroup(string, *models.MulticastMbsGroupMemb) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleQueryAuthSoR(*QueryAuthSoRParams) *models.SorData

	HandleCreateAuthenticationSoR(*CreateAuthenticationSoRParams, *models.SorData)

	HandleCreateEeSubscriptions(string, *models.EeSubscription) (map[string]string, *models.EeSubscription)

	HandleUpdateEesubscriptions(*UpdateEesubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleUpdateEeGroupSubscriptions(*UpdateEeGroupSubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleRemoveEeGroupSubscriptions(*RemoveEeGroupSubscriptionsParams)

	HandleGetOdbData(string) *models.OdbData

	HandleCreateOrUpdateNssaiAck(*CreateOrUpdateNssaiAckParams, *models.NssaiAckData)

	HandleQueryIpSmGwContext(*QueryIpSmGwContextParams) *models.IpSmGwRegistration

	HandleDeleteMessageWaitingData(string)

	HandleGetMultiplePPDataEntries(*GetMultiplePPDataEntriesParams) (*models.PpDataEntryList, *models.ProblemDetails)

	HandleRemoveSmfGroupSubscriptions(*RemoveSmfGroupSubscriptionsParams)

	HandleModifySmfGroupSubscriptions(*ModifySmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryTraceData(*QueryTraceDataParams) (map[string]string, *models.TraceData)

	HandleQuerySmfRegistration(*QuerySmfRegistrationParams) *models.SmfRegistration

	HandleCreateSmsfContext3gpp(string, *models.SmsfRegistration) *models.SmsfRegistration

	HandleCreateSmsfContextNon3gpp(string, *models.SmsfRegistration) *models.SmsfRegistration

	HandleDeleteIpSmGwContext(string)

	HandleGetppData(*GetppDataParams) (map[string]string, *models.PpData, *models.ProblemDetails)

	HandleQueryEEData(*QueryEEDataParams) *models.EeProfileData

	HandleQuerysdmSubscription(*QuerysdmSubscriptionParams)

	HandleQueryAuthenticationStatus(*QueryAuthenticationStatusParams) *models.AuthEvent

	HandleCreateCagUpdateAck(*CreateCagUpdateAckParams, *models.CagAckData)

	HandleQueryProvisionedData(*QueryProvisionedDataParams) *models.ProvisionedDataSets

	HandleCreateMessageWaitingData(string, *models.MessageWaitingData) *models.MessageWaitingData

	HandleDeletePPDataEntry(*DeletePPDataEntryParams) *models.ProblemDetails

	HandleCreateIndividualAuthenticationStatus(*CreateIndividualAuthenticationStatusParams, *models.AuthEvent)

	HandleCreateOperSpecData(*CreateOperSpecDataParams, *map[string]models.OperatorSpecificDataContainer) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleGetHssSubscriptionInfo(*GetHssSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleCreateEeGroupSubscriptions(string, *models.EeSubscription) (map[string]string, *models.EeSubscription)

	HandleQueryUeLocation(*QueryUeLocationParams) *models.LocationInfo

	HandleCreateIpSmGwContext(string, *models.IpSmGwRegistration)

	HandleRemoveSmfSubscriptionsInfo(*RemoveSmfSubscriptionsInfoParams)

	HandleQuerysdmsubscriptions(*QuerysdmsubscriptionsParams) *[]models.SdmSubscription

	HandleModifyHssSDMSubscriptionInfo(*ModifyHssSDMSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetSharedData(*GetSharedDataParams) (*[]models.SharedData, *models.ProblemDetails)

	HandleDelete5GVnGroup(string)

	HandleGetServiceSpecificAuthorizationInfo(*GetServiceSpecificAuthorizationInfoParams) *models.ServiceSpecificAuthorizationInfo

	HandleCreateAuthenticationUPU(*CreateAuthenticationUPUParams, *models.UpuData)

	HandleRemoveeeSubscriptions(*RemoveeeSubscriptionsParams)

	HandleGetSSAuData(*GetSSAuDataParams) *models.ProblemDetails

	HandleQuery5GMbsGroupPPData(*Query5GMbsGroupPPDataParams) (*models.Pp5gMbsGroupProfileData, *models.ProblemDetails)

	HandleDeleteIndividualAuthenticationStatus(*DeleteIndividualAuthenticationStatusParams)

	HandleGetNiddAuData(*GetNiddAuDataParams) *models.ProblemDetails

	HandleQueryLcsBcaData(*QueryLcsBcaDataParams) (map[string]string, *models.LcsBroadcastAssistanceTypesData)

	HandleQuery5mbsData(*Query5mbsDataParams) (map[string]string, *models.MbsSubscriptionData)

	HandleUpdateSmfContext(*UpdateSmfContextParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModify5GVnGroup(*Modify5GVnGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryPeiInformation(string) *models.PeiUpdateInfo

	HandleCreateOrUpdatePeiInformation(string, *models.PeiUpdateInfo) *models.PeiUpdateInfo

	HandleGetSmfGroupSubscriptions(*GetSmfGroupSubscriptionsParams) *models.SmfSubscriptionInfo

	HandleQueryAuthSubsData(*QueryAuthSubsDataParams) *models.AuthenticationSubscription

	HandleQueryPorseData(*QueryPorseDataParams) (map[string]string, *models.ProseSubscriptionData)

	HandleQuerySubsToNotify(*QuerySubsToNotifyParams) *[]models.SubscriptionDataSubscriptions

	HandleQuerySmData(*QuerySmDataParams) (map[string]string, *models.SmSubsData)

	HandleDeleteSmfRegistration(*DeleteSmfRegistrationParams)

	HandleModifyIpSmGwContext(string, *[]models.PatchItem) *models.ProblemDetails

	HandleModifyMessageWaitingData(string, *[]models.PatchItem) *models.ProblemDetails

	HandleCreatePPDataEntry(*CreatePPDataEntryParams, *models.PpDataEntry) (*models.PpDataEntry, *models.ProblemDetails)

	HandleGetSmfSubscriptionInfo(*GetSmfSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleCreateSMFSubscriptions(*CreateSMFSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleCreateAmfContext3gpp(string, *models.Amf3GppAccessRegistration) (map[string]string, *models.Amf3GppAccessRegistration)

	HandleModifyPpData(*ModifyPpDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryGroupEEData(*QueryGroupEEDataParams) *models.EeGroupProfileData

	HandleQuery5GVNGroupPPData(*Query5GVNGroupPPDataParams) *models.Pp5gVnGroupProfileData
}
