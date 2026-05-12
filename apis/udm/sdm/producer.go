/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:37 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package sdm

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"io"
)

func OnGetSmfSelData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSmfSelDataParams

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
	headers, rsp, prob := prod.HandleGetSmfSelData(ctx.Context(), &params)

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
	headers, rsp, prob := prod.HandleGetSharedData(ctx.Context(), &params)

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
	headers, rsp, prob := prod.HandleSubscribe(ctx.Context(), ueId, body)

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
	prob := prod.HandleSorAckInfo(ctx.Context(), supi, body)

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
	prob := prod.HandleUpuAck(ctx.Context(), supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

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
	headers, rsp, prob := prod.HandleGetNSSAI(ctx.Context(), &params)

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
	rsp, prob := prod.HandleGetUeCtxInSmfData(ctx.Context(), &params)

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
	rsp, prob := prod.HandleGetUeCtxInSmsfData(ctx.Context(), &params)

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

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetTraceConfigData(ctx.Context(), &params)

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

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetSmData(ctx.Context(), &params)

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
	headers, rsp, prob := prod.HandleGetLcsMoData(ctx.Context(), &params)

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
	rsp, prob := prod.HandleGetUeCtxInAmfData(ctx.Context(), &params)

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

func OnGetLcsBcaData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetLcsBcaDataParams

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

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetLcsBcaData(ctx.Context(), &params)

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
	headers, rsp, prob := prod.HandleGetV2xData(ctx.Context(), &params)

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
	rsp, prob := prod.HandleUpdateSORInfo(ctx.Context(), supi, body)

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

func OnModifySharedDataSubs(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySharedDataSubsParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
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
	rsp, prob := prod.HandleModifySharedDataSubs(ctx.Context(), &params, body)

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

func OnGetIndividualSharedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetIndividualSharedDataParams

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetIndividualSharedData(ctx.Context(), &params)

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

func OnGetMultipleIdentifiers(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetMultipleIdentifiersParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

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

	// call application handler
	headers, rsp, prob := prod.HandleGetMultipleIdentifiers(ctx.Context(), &params)

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
	prob := prod.HandleSNSSAIsAck(ctx.Context(), supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetEcrData(ctx sbi.RequestContext, prod Producer) {
	var params GetEcrDataParams

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
	headers, rsp, prob := prod.HandleGetEcrData(ctx.Context(), &params)

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

func OnGetMbsData(ctx sbi.RequestContext, prod Producer) {
	var params GetMbsDataParams

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
	headers, rsp, prob := prod.HandleGetMbsData(ctx.Context(), &params)

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

	// read 'uc-purpose'
	params.UcPurpose = ctx.Param("uc-purpose")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetUcData(ctx.Context(), &params)

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
	headers, rsp, prob := prod.HandleSubscribeToSharedData(ctx.Context(), body)

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

func OnGetAmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetAmDataParams

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

	// call application handler
	headers, rsp, prob := prod.HandleGetAmData(ctx.Context(), &params)

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
	headers, rsp, prob := prod.HandleGetSmsData(ctx.Context(), &params)

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
	headers, rsp, prob := prod.HandleGetLcsPrivacyData(ctx.Context(), &params)

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

func OnGetProseData(ctx sbi.RequestContext, prod Producer) {
	var params GetProseDataParams

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
	headers, rsp, prob := prod.HandleGetProseData(ctx.Context(), &params)

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

func OnUnsubscribe(ctx sbi.RequestContext, prod Producer) {
	var params UnsubscribeParams

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribe(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModify(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyParams

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
	rsp, prob := prod.HandleModify(ctx.Context(), &params, body)

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

func OnGetSupiOrGpsi(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSupiOrGpsiParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'mtc-provider-info'
	params.MtcProviderInfo = ctx.Param("mtc-provider-info")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

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

	// read 'af-service-id'
	params.AfServiceId = ctx.Param("af-service-id")

	// read 'requested-gpsi-type'
	params.RequestedGpsiType = ctx.Param("requested-gpsi-type")

	// call application handler
	headers, rsp, prob := prod.HandleGetSupiOrGpsi(ctx.Context(), &params)

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

func OnUnsubscribeForSharedData(ctx sbi.RequestContext, prod Producer) {

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subscriptionId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribeForSharedData(ctx.Context(), subscriptionId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSmsMngtData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSmsMngtDataParams

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
	headers, rsp, prob := prod.HandleGetSmsMngtData(ctx.Context(), &params)

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
	prob := prod.HandleCAGAck(ctx.Context(), supi, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetGroupIdentifiers(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetGroupIdentifiersParams

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

	// call application handler
	headers, rsp, prob := prod.HandleGetGroupIdentifiers(ctx.Context(), &params)

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

func OnGetDataSets(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetDataSetsParams

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

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetDataSets(ctx.Context(), &params)

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
	HandleGetSmfSelData(context.Context, *GetSmfSelDataParams) (map[string]string, *models.SmfSelectionSubscriptionData, *models.ProblemDetails)

	HandleGetSharedData(context.Context, *GetSharedDataParams) (map[string]string, *[]models.SharedData, *models.ProblemDetails)

	HandleSubscribe(context.Context, string, *models.SdmSubscription) (map[string]string, *models.SdmSubscription, *models.ProblemDetails)

	HandleSorAckInfo(context.Context, string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleUpuAck(context.Context, string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetNSSAI(context.Context, *GetNSSAIParams) (map[string]string, *models.Nssai, *models.ProblemDetails)

	HandleGetUeCtxInSmfData(context.Context, *GetUeCtxInSmfDataParams) (*models.UeContextInSmfData, *models.ProblemDetails)

	HandleGetUeCtxInSmsfData(context.Context, *GetUeCtxInSmsfDataParams) (*models.UeContextInSmsfData, *models.ProblemDetails)

	HandleGetTraceConfigData(context.Context, *GetTraceConfigDataParams) (map[string]string, *models.TraceDataResponse, *models.ProblemDetails)

	HandleGetSmData(context.Context, *GetSmDataParams) (map[string]string, *models.SmSubsData, *models.ProblemDetails)

	HandleGetLcsMoData(context.Context, *GetLcsMoDataParams) (map[string]string, *models.LcsMoData, *models.ProblemDetails)

	HandleGetUeCtxInAmfData(context.Context, *GetUeCtxInAmfDataParams) (*models.UeContextInAmfData, *models.ProblemDetails)

	HandleGetLcsBcaData(context.Context, *GetLcsBcaDataParams) (map[string]string, *models.LcsBroadcastAssistanceTypesData, *models.ProblemDetails)

	HandleGetV2xData(context.Context, *GetV2xDataParams) (map[string]string, *models.V2xSubscriptionData, *models.ProblemDetails)

	HandleUpdateSORInfo(context.Context, string, *models.SorUpdateInfo) (*models.SorInfo, *models.ProblemDetails)

	HandleModifySharedDataSubs(context.Context, *ModifySharedDataSubsParams, *models.SdmSubsModification) (*models.Schema, *models.ProblemDetails)

	HandleGetIndividualSharedData(context.Context, *GetIndividualSharedDataParams) (map[string]string, *models.SharedData, *models.ProblemDetails)

	HandleGetMultipleIdentifiers(context.Context, *GetMultipleIdentifiersParams) (map[string]string, *map[string]models.SupiInfo, *models.ProblemDetails)

	HandleSNSSAIsAck(context.Context, string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetEcrData(context.Context, *GetEcrDataParams) (map[string]string, *models.EnhancedCoverageRestrictionData, *models.ProblemDetails)

	HandleGetMbsData(context.Context, *GetMbsDataParams) (map[string]string, *models.MbsSubscriptionData, *models.ProblemDetails)

	HandleGetUcData(context.Context, *GetUcDataParams) (map[string]string, *models.UcSubscriptionData, *models.ProblemDetails)

	HandleSubscribeToSharedData(context.Context, *models.SdmSubscription) (map[string]string, *models.SdmSubscription, *models.ProblemDetails)

	HandleGetAmData(context.Context, *GetAmDataParams) (map[string]string, *models.AccessAndMobilitySubscriptionData, *models.ProblemDetails)

	HandleGetSmsData(context.Context, *GetSmsDataParams) (map[string]string, *models.SmsSubscriptionData, *models.ProblemDetails)

	HandleGetLcsPrivacyData(context.Context, *GetLcsPrivacyDataParams) (map[string]string, *models.LcsPrivacyData, *models.ProblemDetails)

	HandleGetProseData(context.Context, *GetProseDataParams) (map[string]string, *models.ProseSubscriptionData, *models.ProblemDetails)

	HandleUnsubscribe(context.Context, *UnsubscribeParams) *models.ProblemDetails

	HandleModify(context.Context, *ModifyParams, *models.SdmSubsModification) (*models.Schema, *models.ProblemDetails)

	HandleGetSupiOrGpsi(context.Context, *GetSupiOrGpsiParams) (map[string]string, *models.IdTranslationResult, *models.ProblemDetails)

	HandleUnsubscribeForSharedData(context.Context, string) *models.ProblemDetails

	HandleGetSmsMngtData(context.Context, *GetSmsMngtDataParams) (map[string]string, *models.SmsManagementSubscriptionData, *models.ProblemDetails)

	HandleCAGAck(context.Context, string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetGroupIdentifiers(context.Context, *GetGroupIdentifiersParams) (map[string]string, *models.GroupIdentifiers, *models.ProblemDetails)

	HandleGetDataSets(context.Context, *GetDataSetsParams) (map[string]string, *models.SubscriptionDataSets, *models.ProblemDetails)
}
