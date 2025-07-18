/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"io"
)

func OnUnsubscribeForSharedData(ctx sbi.RequestContext, prod Producer) {

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribeForSharedData(subscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetMultipleIdentifiers(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetMultipleIdentifiersParams

	// read 'gpsi-list'
	gpsiListStr := ctx.Param("gpsi-list")
	if len(gpsiListStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "gpsi-list is required"), nil)
		return
	}

	if params.GpsiList, err = models.ArrayOfStringFromString(gpsiListStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsi-list failed: %+v", err)), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp, prob := prod.HandleGetMultipleIdentifiers(&params)

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

func OnGetNSSAI(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetNSSAIParams

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)), nil)
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetNSSAI(&params)

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

func OnGetLcsMoData(ctx sbi.RequestContext, prod Producer) {
	var params GetLcsMoDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetLcsMoData(&params)

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

func OnCAGAck(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.AcknowledgeInfo
	body = new(models.AcknowledgeInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	prob := prod.HandleCAGAck(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetIndividualSharedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetIndividualSharedDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'sharedDataId'
	sharedDataIdStr := ctx.Param("sharedDataId")
	if len(sharedDataIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sharedDataId is required"), nil)
		return
	}

	if params.SharedDataId, err = models.ArrayOfStringFromString(sharedDataIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse sharedDataId failed: %+v", err)), nil)
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

func OnGetSmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSmDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.SnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
			return
		}
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetSmData(&params)

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

func OnGetV2xData(ctx sbi.RequestContext, prod Producer) {
	var params GetV2xDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp, prob := prod.HandleGetV2xData(&params)

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

func OnGetSmsData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSmsDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp, prob := prod.HandleGetSmsData(&params)

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

func OnSubscribe(ctx sbi.RequestContext, prod Producer) {
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
	headers, rsp, prob := prod.HandleSubscribe(ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnUnsubscribe(ctx sbi.RequestContext, prod Producer) {
	var params UnsubscribeParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribe(&params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpuAck(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.AcknowledgeInfo
	body = new(models.AcknowledgeInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	prob := prod.HandleUpuAck(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetDataSets(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetDataSetsParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dataset-names is required"), nil)
		return
	}

	if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)), nil)
		return
	}

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdNidFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)), nil)
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetDataSets(&params)

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

func OnGetSmfSelData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSmfSelDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)), nil)
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetSmfSelData(&params)

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

func OnGetUeCtxInSmfData(ctx sbi.RequestContext, prod Producer) {
	var params GetUeCtxInSmfDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInSmfData(&params)

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

func OnGetTraceConfigData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetTraceConfigDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp, prob := prod.HandleGetTraceConfigData(&params)

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

func OnGetSmsMngtData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSmsMngtDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetSmsMngtData(&params)

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

func OnSorAckInfo(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.AcknowledgeInfo
	body = new(models.AcknowledgeInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	prob := prod.HandleSorAckInfo(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifySharedDataSubs(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySharedDataSubsParams

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SdmSubsModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySharedDataSubs(&params, body)

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

func OnGetUeCtxInAmfData(ctx sbi.RequestContext, prod Producer) {
	var params GetUeCtxInAmfDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInAmfData(&params)

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

func OnGetAmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetAmDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdNidFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'adjacent-plmns'
	adjacentPlmnsStr := ctx.Param("adjacent-plmns")
	if len(adjacentPlmnsStr) > 0 {
		if params.AdjacentPlmns, err = models.ArrayOfPlmnIdFromString(adjacentPlmnsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse adjacent-plmns failed: %+v", err)), nil)
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)), nil)
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetAmData(&params)

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

func OnModify(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyParams

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

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
	body := new(models.SdmSubsModification)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify(&params, body)

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

func OnSNSSAIsAck(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.AcknowledgeInfo
	body = new(models.AcknowledgeInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	prob := prod.HandleSNSSAIsAck(supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetProseData(ctx sbi.RequestContext, prod Producer) {
	var params GetProseDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetProseData(&params)

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

func OnGetUcData(ctx sbi.RequestContext, prod Producer) {
	var params GetUcDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'uc-purpose'
	params.UcPurpose = ctx.Param("uc-purpose")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetUcData(&params)

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

func OnGetSupiOrGpsi(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSupiOrGpsiParams

	// read 'af-service-id'
	params.AfServiceId = ctx.Param("af-service-id")

	// read 'requested-gpsi-type'
	params.RequestedGpsiType = ctx.Param("requested-gpsi-type")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

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

	// read 'mtc-provider-info'
	params.MtcProviderInfo = ctx.Param("mtc-provider-info")

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetSupiOrGpsi(&params)

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

func OnUpdateSORInfo(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.SorUpdateInfo
	body = new(models.SorUpdateInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, prob := prod.HandleUpdateSORInfo(supi, body)

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

func OnSubscribeToSharedData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SdmSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleSubscribeToSharedData(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnGetEcrData(ctx sbi.RequestContext, prod Producer) {
	var params GetEcrDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp, prob := prod.HandleGetEcrData(&params)

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

func OnGetUeCtxInSmsfData(ctx sbi.RequestContext, prod Producer) {
	var params GetUeCtxInSmsfDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInSmsfData(&params)

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

func OnGetMbsData(ctx sbi.RequestContext, prod Producer) {
	var params GetMbsDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetMbsData(&params)

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

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetSharedData(&params)

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

func OnGetGroupIdentifiers(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetGroupIdentifiersParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

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

	// call application handler
	headers, rsp, prob := prod.HandleGetGroupIdentifiers(&params)

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

func OnGetLcsPrivacyData(ctx sbi.RequestContext, prod Producer) {
	var params GetLcsPrivacyDataParams

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
	headers, rsp, prob := prod.HandleGetLcsPrivacyData(&params)

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

func OnGetLcsBcaData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetLcsBcaDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp, prob := prod.HandleGetLcsBcaData(&params)

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

type Producer interface {
	HandleUnsubscribeForSharedData(string) *models.ProblemDetails

	HandleGetMultipleIdentifiers(*GetMultipleIdentifiersParams) (map[string]string, *map[string]models.SupiInfo, *models.ProblemDetails)

	HandleGetNSSAI(*GetNSSAIParams) (map[string]string, *models.Nssai, *models.ProblemDetails)

	HandleGetLcsMoData(*GetLcsMoDataParams) (map[string]string, *models.LcsMoData, *models.ProblemDetails)

	HandleCAGAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetIndividualSharedData(*GetIndividualSharedDataParams) (map[string]string, *models.SharedData, *models.ProblemDetails)

	HandleGetSmData(*GetSmDataParams) (map[string]string, *models.SmSubsData, *models.ProblemDetails)

	HandleGetV2xData(*GetV2xDataParams) (map[string]string, *models.V2xSubscriptionData, *models.ProblemDetails)

	HandleGetSmsData(*GetSmsDataParams) (map[string]string, *models.SmsSubscriptionData, *models.ProblemDetails)

	HandleSubscribe(string, *models.SdmSubscription) (map[string]string, *models.SdmSubscription, *models.ProblemDetails)

	HandleUnsubscribe(*UnsubscribeParams) *models.ProblemDetails

	HandleUpuAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetDataSets(*GetDataSetsParams) (map[string]string, *models.SubscriptionDataSets, *models.ProblemDetails)

	HandleGetSmfSelData(*GetSmfSelDataParams) (map[string]string, *models.SmfSelectionSubscriptionData, *models.ProblemDetails)

	HandleGetUeCtxInSmfData(*GetUeCtxInSmfDataParams) (*models.UeContextInSmfData, *models.ProblemDetails)

	HandleGetTraceConfigData(*GetTraceConfigDataParams) (map[string]string, *models.TraceDataResponse, *models.ProblemDetails)

	HandleGetSmsMngtData(*GetSmsMngtDataParams) (map[string]string, *models.SmsManagementSubscriptionData, *models.ProblemDetails)

	HandleSorAckInfo(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleModifySharedDataSubs(*ModifySharedDataSubsParams, *models.SdmSubsModification) (*models.Schema, *models.ProblemDetails)

	HandleGetUeCtxInAmfData(*GetUeCtxInAmfDataParams) (*models.UeContextInAmfData, *models.ProblemDetails)

	HandleGetAmData(*GetAmDataParams) (map[string]string, *models.AccessAndMobilitySubscriptionData, *models.ProblemDetails)

	HandleModify(*ModifyParams, *models.SdmSubsModification) (*models.Schema, *models.ProblemDetails)

	HandleSNSSAIsAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetProseData(*GetProseDataParams) (map[string]string, *models.ProseSubscriptionData, *models.ProblemDetails)

	HandleGetUcData(*GetUcDataParams) (map[string]string, *models.UcSubscriptionData, *models.ProblemDetails)

	HandleGetSupiOrGpsi(*GetSupiOrGpsiParams) (map[string]string, *models.IdTranslationResult, *models.ProblemDetails)

	HandleUpdateSORInfo(string, *models.SorUpdateInfo) (*models.SorInfo, *models.ProblemDetails)

	HandleSubscribeToSharedData(*models.SdmSubscription) (map[string]string, *models.SdmSubscription, *models.ProblemDetails)

	HandleGetEcrData(*GetEcrDataParams) (map[string]string, *models.EnhancedCoverageRestrictionData, *models.ProblemDetails)

	HandleGetUeCtxInSmsfData(*GetUeCtxInSmsfDataParams) (*models.UeContextInSmsfData, *models.ProblemDetails)

	HandleGetMbsData(*GetMbsDataParams) (map[string]string, *models.MbsSubscriptionData, *models.ProblemDetails)

	HandleGetSharedData(*GetSharedDataParams) (map[string]string, *[]models.SharedData, *models.ProblemDetails)

	HandleGetGroupIdentifiers(*GetGroupIdentifiersParams) (map[string]string, *models.GroupIdentifiers, *models.ProblemDetails)

	HandleGetLcsPrivacyData(*GetLcsPrivacyDataParams) (map[string]string, *models.LcsPrivacyData, *models.ProblemDetails)

	HandleGetLcsBcaData(*GetLcsBcaDataParams) (map[string]string, *models.LcsBroadcastAssistanceTypesData, *models.ProblemDetails)
}
