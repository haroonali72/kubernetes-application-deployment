package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"kubernetes-services-deployment/core"
	"kubernetes-services-deployment/types"
	"kubernetes-services-deployment/utils"
	"net/http"
)

// @host engine.swagger.io
// @BasePath /api/v1/

// @Summary deploy services on kubernetes cluster
// @Description deploy services on kubernetes cluster

// @Accept  json
// @Produce  json
// @router /api/v1/registry [post]
func (c *KubeController) CreateRegistrySecret(g *gin.Context) {
	var req types.RegistryRequest
	err := g.ShouldBind(&req)
	if err != nil {
		utils.Error.Println(err)
		utils.NewError(g, http.StatusBadRequest, err)
		return
	}
	kubeClient, err := core.GetKubernetesClient(req.ProjectId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": err.Error()})
		return
	}
	data, err := kubeClient.CreateDockerRegistryCredentials(&req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "service secrets credentials creation failed.", "Error": err.Error()})
		return
	}
	d, err := json.Marshal(data)
	if err != nil {
		utils.Error.Println(err)
		g.JSON(http.StatusOK, gin.H{"status": "service secrets created successfully", "error": nil, "data": ""})
		return
	}
	g.JSON(http.StatusOK, gin.H{"status": "service secrets created successfully", "error": nil, "data": string(d)})

}

// @host engine.swagger.io
// @BasePath /api/v1/

// @Summary deploy services on kubernetes cluster
// @Description deploy services on kubernetes cluster
// @Param project_id header string	true "project id"
// @Param name path string true "Name of the kubernetes service"
// @Param namespace path string true "Namespace of the kubernetes service"
// @Accept  json
// @Produce  json
// @router /api/v1/registry/{namespace}/{name} [get]
func (c *KubeController) GetRegistrySecret(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	projectId := g.GetHeader("project_id")

	if projectId == "" {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": "project_id is missing in request"})
		return
	}
	if name == "" {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": "service name is not invalid"})
		return
	}
	kubeClient, err := core.GetKubernetesClient(&projectId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": err.Error()})
		return
	}

	data, err := kubeClient.GetDockerRegistryCredentials(name, namespace)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": err.Error()})
		return
	}
	d, err := json.Marshal(data)
	if err != nil {
		utils.Error.Println(err)
		g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "data": ""})
		return
	}
	g.JSON(http.StatusOK, gin.H{"error": nil, "data": string(d)})
}

// @host engine.swagger.io
// @BasePath /api/v1/

// @Summary deploy services on kubernetes cluster
// @Description deploy services on kubernetes cluster
// @Param project_id header string	true "project id"
// @Param name path string true "Name of the kubernetes service"
// @Param namespace path string true "Namespace of the kubernetes service"
// @Accept  json
// @Produce  json
// @router /api/v1/registry/{namespace}/{name} [delete]
func (c *KubeController) DeleteRegistrySecret(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	projectId := g.GetHeader("project_id")

	if projectId == "" {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": "project_id is missing in request"})
		return
	}
	kubeClient, err := core.GetKubernetesClient(&projectId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": err.Error()})
		return
	}
	if name == "" {
		g.JSON(http.StatusInternalServerError, gin.H{"data": "", "Error": "service name is not invalid"})
		return
	}
	err = kubeClient.DeleteDockerRegistryCredentials(name, namespace)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed to delete secrets", "Error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"error": nil, "status": "secrets deleted successfully"})
}
