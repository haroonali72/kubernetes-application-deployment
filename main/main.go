package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"kubernetes-services-deployment/constants"
	"kubernetes-services-deployment/controllers"
	_ "kubernetes-services-deployment/controllers/docs"
	"kubernetes-services-deployment/utils"
	"os"
	"time"
)

func init() {
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /ksd/api/v1

func main() {

	e := gin.New()
	utils.InitFlags()
	constants.CacheObj = cache.New(5*time.Minute, 5*time.Minute)
	if constants.ServicePort == "" {
		constants.ServicePort = "8089"
	}
	c, _ := controllers.NewController()
	v1 := e.Group("/ksd/api/v1")
	{
		/*dag := v1.Group("/kubernetes")
		{
			dag.POST("deploy", c.DeployService)
		}*/
		v1.POST("/solution", c.DeploySolution)
		v1.GET("/solution", c.GetSolution)
		v1.DELETE("/solution", c.DeleteSolution)
		v1.PATCH("/solution", c.PatchSolution)
		v1.PUT("/solution", c.PutSolution)
		///statefulsets APIs
		v1.GET("/statefulsets/:namespace", c.ListStatefulSetsStatus)
		v1.GET("/statefulsets/:namespace/:name", c.GetStatefulSetsStatus)
		v1.DELETE("/statefulsets/:namespace/:name", c.DeleteStatefulSetsStatus)

		//secrets APIs
		v1.POST("registry", c.CreateRegistrySecret)
		v1.GET("/registry/:namespace/:name", c.GetStatefulSetsStatus)
		v1.DELETE("/registry/:namespace/:name", c.DeleteRegistrySecret)
		//deployment APIs
		v1.GET("/deployment/:namespace", c.ListDeploymentStatus)
		v1.GET("/deployment/:namespace/:name", c.GetDeploymentStatus)
		v1.DELETE("/deployment/:namespace/:name", c.DeleteDeployment)

		v1.GET("/kubeservice/:namespace", c.ListKubernetesServices)
		v1.GET("/kubeservice/:namespace/:name", c.GetKubernetesService)
		v1.DELETE("/kubeservice/:namespace/:name", c.DeleteKubernetesService)

		v1.GET("/kubeservice/:namespace/:name/endpoint", c.GetKubernetesServiceExternalIp)

	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.Run(":" + constants.ServicePort)
	//e.Logger.Fatal(e.Start(":" + constants.ServicePort))

}

/*func main_X() {

	client, err := kubernetes.NewForConfig(&rest.Config{Host: "https://54.237.228.34:6443", Username: "cloudplex", Password: "64bdySICej", TLSClientConfig: rest.TLSClientConfig{Insecure: true}})
	utils.Error.Println(err)
	pods, err := client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := range pods.Items {
		utils.Info.Println(pods.Items[i].Name, pods.Items[i].Namespace)
	}

}*/
