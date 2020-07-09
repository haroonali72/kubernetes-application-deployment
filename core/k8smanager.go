package core

import (
	"bitbucket.org/cloudplex-devs/kubernetes-services-deployment/constants"
	pb "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	v1alpha "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/kubernetes-custom-apis/core/v1"
	"bitbucket.org/cloudplex-devs/kubernetes-services-deployment/utils"
	"context"
	"encoding/json"
	"errors"
	"reflect"
)

type Server struct {
}

func (s *Server) CreateService(ctx context.Context, request *pb.ServiceRequest) (response *pb.SerivceFResponse, err error) {
	response = new(pb.SerivceFResponse)
	utils.Info.Println(reflect.TypeOf(ctx))
	cpCtx := &Context{}
	cpCtx.Keys = make(map[string]interface{})
	cpCtx.Keys[constants.AuthTokenKey] = request.Token
	uId, CID, err := utils.GetUserIDCompanyID(request.Token)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	if request.ProjectId == "" || request.CompanyId == "" {
		utils.Error.Println("ProjectID or CompanyID is empty")
		return response, errors.New("ProjectID or CompanyID is empty")
	}
	cpCtx.Keys["company_id"] = CID
	cpCtx.Keys["user"] = uId
	cpCtx.Keys["project_id"] = request.ProjectId
	agent, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	agent.CpCtx = cpCtx
	agent.CompanyId = request.CompanyId
	agent.ProjectId = request.ProjectId

	runtimeObj := v1alpha.RuntimeConfig{}
	err = json.Unmarshal(request.Service, &runtimeObj)
	if err != nil {
		return response, err
	}

	responseObj, err := agent.crdManager(runtimeObj, string(constants.POST))
	//service, err := agent.AgentCrdManager(constants.POST, request)
	if err != nil {
		cpCtx.SendFrontendLogs(err.Error(), constants.LOGGING_LEVEL_ERROR)
		return response, err
	}
	cpCtx.SendFrontendLogs(responseObj, constants.LOGGING_LEVEL_INFO)
	if responseObj.Error != "" {
		cpCtx.SendFrontendLogs(responseObj.Error, constants.LOGGING_LEVEL_ERROR)
		return response, errors.New(responseObj.Error)
	}

	raw, err := json.Marshal(responseObj.Data)
	if err != nil {
		return response, err
	}
	response.Service = raw
	return response, nil
}
func (s *Server) GetService(ctx context.Context, request *pb.ServiceRequest) (response *pb.SerivceFResponse, err error) {
	response = new(pb.SerivceFResponse)
	utils.Info.Println(reflect.TypeOf(ctx))
	cpCtx := &Context{}
	cpCtx.Keys = make(map[string]interface{})
	cpCtx.Keys[constants.AuthTokenKey] = request.Token
	uId, CID, err := utils.GetUserIDCompanyID(request.Token)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	if request.ProjectId == "" || request.CompanyId == "" {
		utils.Error.Println("ProjectID or CompanyID is empty")
		return response, errors.New("ProjectID or CompanyID is empty")
	}
	cpCtx.Keys["company_id"] = CID
	cpCtx.Keys["user"] = uId
	cpCtx.Keys["project_id"] = request.ProjectId
	agent, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	agent.CpCtx = cpCtx
	agent.CompanyId = request.CompanyId
	agent.ProjectId = request.ProjectId

	runtimeObj := v1alpha.RuntimeConfig{}
	err = json.Unmarshal(request.Service, &runtimeObj)
	if err != nil {
		return response, err
	}

	responseObj, err := agent.crdManager(runtimeObj, string(constants.GET))
	//service, err := agent.AgentCrdManager(constants.POST, request)
	if err != nil {
		cpCtx.SendFrontendLogs(err.Error(), constants.LOGGING_LEVEL_ERROR)
		return response, err
	}
	cpCtx.SendFrontendLogs(responseObj, constants.LOGGING_LEVEL_INFO)
	if responseObj.Error != "" {
		cpCtx.SendFrontendLogs(responseObj.Error, constants.LOGGING_LEVEL_ERROR)
	}

	//raw, err := json.Marshal(responseObj.Data)
	//if err != nil {
	//	return response, err
	//}
	response.Service = []byte(responseObj.Data)
	response.PodErrors = responseObj.PodErrors

	/*conn, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	service, err := conn.AgentCrdManager(constants.GET, request)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	utils.Info.Println(string(service))
	response.Service = service*/
	return response, nil
}
func (s *Server) DeleteService(ctx context.Context, request *pb.ServiceRequest) (response *pb.SerivceFResponse, err error) {
	response = new(pb.SerivceFResponse)
	utils.Info.Println(reflect.TypeOf(ctx))
	cpCtx := &Context{}
	cpCtx.Keys = make(map[string]interface{})
	cpCtx.Keys[constants.AuthTokenKey] = request.Token
	uId, CID, err := utils.GetUserIDCompanyID(request.Token)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	if request.ProjectId == "" || request.CompanyId == "" {
		utils.Error.Println("ProjectID or CompanyID is empty")
		return response, errors.New("ProjectID or CompanyID is empty")
	}
	cpCtx.Keys["company_id"] = CID
	cpCtx.Keys["user"] = uId
	cpCtx.Keys["project_id"] = request.ProjectId
	agent, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	agent.CpCtx = cpCtx
	agent.CompanyId = request.CompanyId
	agent.ProjectId = request.ProjectId
	runtimeObj := v1alpha.RuntimeConfig{}
	err = json.Unmarshal(request.Service, &runtimeObj)
	if err != nil {
		return response, err
	}

	responseObj, err := agent.crdManager(runtimeObj, string(constants.DELETE))
	//service, err := agent.AgentCrdManager(constants.POST, request)
	if err != nil {
		cpCtx.SendFrontendLogs(err.Error(), constants.LOGGING_LEVEL_ERROR)
		return response, err
	}
	cpCtx.SendFrontendLogs(responseObj, constants.LOGGING_LEVEL_INFO)
	if responseObj.Error != "" {
		cpCtx.SendFrontendLogs(responseObj.Error, constants.LOGGING_LEVEL_ERROR)
	}

	raw, err := json.Marshal(responseObj.Data)
	if err != nil {
		return response, err
	}
	response.Service = raw
	/*conn, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	service, err := conn.AgentCrdManager(constants.DELETE, request)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	utils.Info.Println(string(service))
	response.Service = service*/
	return response, nil
}
func (s *Server) PatchService(ctx context.Context, request *pb.ServiceRequest) (response *pb.SerivceFResponse, err error) {
	response = new(pb.SerivceFResponse)
	utils.Info.Println(reflect.TypeOf(ctx))
	cpCtx := &Context{}
	cpCtx.Keys = make(map[string]interface{})
	cpCtx.Keys[constants.AuthTokenKey] = request.Token
	uId, CID, err := utils.GetUserIDCompanyID(request.Token)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	if request.ProjectId == "" || request.CompanyId == "" {
		utils.Error.Println("ProjectID or CompanyID is empty")
		return response, errors.New("ProjectID or CompanyID is empty")
	}
	cpCtx.Keys["company_id"] = CID
	cpCtx.Keys["user"] = uId
	cpCtx.Keys["project_id"] = request.ProjectId
	agent, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	agent.CpCtx = cpCtx
	agent.CompanyId = request.CompanyId
	agent.ProjectId = request.ProjectId

	runtimeObj := v1alpha.RuntimeConfig{}
	err = json.Unmarshal(request.Service, &runtimeObj)
	if err != nil {
		return response, err
	}

	responseObj, err := agent.crdManager(runtimeObj, string(constants.PATCH))
	//service, err := agent.AgentCrdManager(constants.POST, request)
	if err != nil {
		cpCtx.SendFrontendLogs(err.Error(), constants.LOGGING_LEVEL_ERROR)
		return response, err
	}
	cpCtx.SendFrontendLogs(responseObj, constants.LOGGING_LEVEL_INFO)
	if responseObj.Error != "" {
		cpCtx.SendFrontendLogs(responseObj.Error, constants.LOGGING_LEVEL_ERROR)
	}
	raw, err := json.Marshal(responseObj.Data)
	if err != nil {
		return response, err
	}
	response.Service = raw
	/*conn, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	service, err := conn.AgentCrdManager(constants.PATCH, request)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	utils.Info.Println(string(service))
	response.Service = service*/
	return response, nil
}
func (s *Server) PutService(ctx context.Context, request *pb.ServiceRequest) (response *pb.SerivceFResponse, err error) {
	response = new(pb.SerivceFResponse)
	utils.Info.Println(reflect.TypeOf(ctx))
	cpCtx := &Context{}
	cpCtx.Keys = make(map[string]interface{})
	cpCtx.Keys[constants.AuthTokenKey] = request.Token
	uId, CID, err := utils.GetUserIDCompanyID(request.Token)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	if request.ProjectId == "" || request.CompanyId == "" {
		utils.Error.Println("ProjectID or CompanyID is empty")
		return response, errors.New("ProjectID or CompanyID is empty")
	}
	cpCtx.Keys["company_id"] = CID
	cpCtx.Keys["user"] = uId
	cpCtx.Keys["project_id"] = request.ProjectId
	agent, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}
	agent.CpCtx = cpCtx
	agent.CompanyId = request.CompanyId
	agent.ProjectId = request.ProjectId

	runtimeObj := v1alpha.RuntimeConfig{}
	err = json.Unmarshal(request.Service, &runtimeObj)
	if err != nil {
		return response, err
	}

	responseObj, err := agent.crdManager(runtimeObj, string(constants.PUT))
	//service, err := agent.AgentCrdManager(constants.POST, request)
	if err != nil {
		cpCtx.SendFrontendLogs(err.Error(), constants.LOGGING_LEVEL_ERROR)
		return response, err
	}
	cpCtx.SendFrontendLogs(responseObj, constants.LOGGING_LEVEL_INFO)
	if responseObj.Error != "" {
		cpCtx.SendFrontendLogs(responseObj.Error, constants.LOGGING_LEVEL_ERROR)
	}
	raw, err := json.Marshal(responseObj.Data)
	if err != nil {
		return response, err
	}
	response.Service = raw

	/*conn, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	service, err := conn.AgentCrdManager(constants.PUT, request)
	if err != nil {
		utils.Error.Println(err)
		return response, err
	}

	utils.Info.Println(string(service))
	response.Service = service*/

	return response, nil
}
