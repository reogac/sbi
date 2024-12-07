package upf2upmf

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n42"
	"fmt"
	"net/http"
)

const SERVICE_PATH string = "upf2upmf"

var _routes = sbi.SbiRoutes{
	{
		Label:   "Heartbeat",
		Method:  http.MethodGet,
		Path:    "heartbeat",
		Handler: OnHeartbeat,
	},
	{
		Label:   "Activate",
		Method:  http.MethodGet,
		Path:    "activate",
		Handler: OnActivate,
	},
	{
		Label:   "Activate",
		Method:  http.MethodGet,
		Path:    "deactive",
		Handler: OnDeactivate,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   "upf2upmf",
		Routes:  _routes,
		Handler: p,
	}
}
func Heartbeat(client sbi.ConsumerClient, body n42.HeartbeatRequest) (rsp *n42.HeartbeatResponse, err error) {

	//create a request
	path := fmt.Sprintf("%s/heartbeat", SERVICE_PATH)
	req := sbi.NewRequest(path, http.MethodGet, &body)

	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if response.GetCode() == 200 {
		rsp = &n42.HeartbeatResponse{}
		err = response.DecodeBody(rsp)
	} else {
		prob := &models.ProblemDetails{}
		if err = response.DecodeBody(prob); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.GetStatus(), prob.Detail)
		}
	}
	return
}

func UpfActivate(client sbi.ConsumerClient, body n42.UpfActivateQuery) (rsp *n42.UpfActivate, err error) {
	//create a request
	path := fmt.Sprintf("%s/activate", SERVICE_PATH)
	req := sbi.NewRequest(path, http.MethodGet, &body)
	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	//handle the response
	if response.GetCode() == 200 {
		rsp = &n42.UpfActivate{}
		err = response.DecodeBody(rsp)
	} else {
		prob := &models.ProblemDetails{}
		if err = response.DecodeBody(prob); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.GetStatus(), prob.Detail)
		}
	}
	return
}

func UpfDeactivate(client sbi.ConsumerClient, body n42.UpfDeactivateQuery) (rsp *n42.UpfDeactivate, err error) {
	//create a request
	path := fmt.Sprintf("%s/deactivate", SERVICE_PATH)
	req := sbi.NewRequest(path, http.MethodGet, &body)

	//send the request
	var response *sbi.Response
	if response, err = client.Send(req); err != nil {
		return
	}
	//handle the response
	if response.GetCode() == 200 {
		rsp = &n42.UpfDeactivate{}
		err = response.DecodeBody(rsp)
	} else {
		prob := &models.ProblemDetails{}
		if err = response.DecodeBody(prob); err == nil {
			err = fmt.Errorf("%s: Problem=%s", response.GetStatus(), prob.Detail)
		}
	}
	return
}
