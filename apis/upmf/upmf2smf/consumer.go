package upmf2smf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upmf2smf"

var _routes = sbi.SbiRoutes{
	{
		Label:   "PathQuery",
		Method:  http.MethodGet,
		Path:    "path",
		Handler: OnGetUpfPath,
	},
	{
		Label:   "Subscribe",
		Method:  http.MethodGet,
		Path:    "sub",
		Handler: OnSmfSubscribe,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   SERVICE_PATH,
		Routes:  _routes,
		Handler: p,
	}
}

func Subscribe(client sbi.ConsumerClient, body n43.SmfSubscriptionRequest) (rsp *n43.SmfSubscriptionResponse, err error) {

	//create a request
	path := fmt.Sprintf("%s/sub", SERVICE_PATH)
	req := sbi.NewRequest(path, http.MethodGet, &body)

	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.GetCode() == http.StatusCreated {
		rsp = &n43.SmfSubscriptionResponse{}
		err = response.DecodeBody(rsp)
	} else {
		prob := &models.ProblemDetails{}
		if err = response.DecodeBody(prob); err == nil {
			err = prob
		} else {
			err = fmt.Errorf("%d: %s", response.GetCode(), response.GetStatus())
		}
	}
	return
}
func GetUpfPath(client sbi.ConsumerClient, body n43.UpfPathQuery) (upfPath *n43.UpfPath, err error) {

	//create a request
	path := fmt.Sprintf("%s/path", SERVICE_PATH)
	req := sbi.NewRequest(path, http.MethodGet, &body)
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.GetCode() == http.StatusCreated {
		upfPath = &n43.UpfPath{}
		err = response.DecodeBody(upfPath)
	} else {
		prob := &models.ProblemDetails{}
		if err = response.DecodeBody(prob); err == nil {
			err = prob
		} else {
			err = fmt.Errorf("%d: %s", response.GetCode(), response.GetStatus())
		}
	}
	return
}
