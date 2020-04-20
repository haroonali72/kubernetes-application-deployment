package core

import (
	"context"
	"errors"
	pb "kubernetes-services-deployment/core/proto"
	"kubernetes-services-deployment/utils"
	"reflect"
)

func (s *Server) GetK8SResource(ctx context.Context, request *pb.KubernetesResourceRequest) (response *pb.KubernetesResourceResponse, err error) {
	response = new(pb.KubernetesResourceResponse)
	utils.Info.Println(reflect.TypeOf(ctx))

	if request.CompanyId == "" || request.ProjectId == "" {
		return &pb.KubernetesResourceResponse{}, errors.New("projectId or companyId must not be empty")
	}

	agent, err := GetGrpcAgentConnection()
	if err != nil {
		utils.Error.Println(err)
		return &pb.KubernetesResourceResponse{}, err
	}

	err = agent.InitializeAgentClient(request.ProjectId, request.CompanyId)
	if err != nil {
		utils.Error.Println(err)
		return &pb.KubernetesResourceResponse{}, err
	}

	defer agent.connection.Close()

	resp, err := agent.GetK8sResources(ctx, request)
	if err != nil {
		return &pb.KubernetesResourceResponse{}, err
	}

	response.Resource = resp
	return response, err

}

//func convertToK8sStruct(resourceData []byte) error {
//	decode := scheme.Codecs.UniversalDeserializer().Decode
//	obj, _, err := decode(resourceData, nil, nil)
//	if err != nil {
//		utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
//		return errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
//	}
//	switch o := obj.(type) {
//	case *v2.Deployment:
//		var deployment v2.Deployment
//		err := json.Unmarshal(resourceData, &deployment)
//	case *v2.DaemonSet:
//	case *v2.ReplicaSet:
//
//	}
//
//	return nil
//}