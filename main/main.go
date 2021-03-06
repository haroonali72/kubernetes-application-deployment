package main

import (
	"bitbucket.org/cloudplex-devs/kubernetes-services-deployment/constants"
	"bitbucket.org/cloudplex-devs/kubernetes-services-deployment/controllers"
	"bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core"
	pb "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	_ "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/docs"
	"bitbucket.org/cloudplex-devs/kubernetes-services-deployment/types"
	"bitbucket.org/cloudplex-devs/kubernetes-services-deployment/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func init() {
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

// @title Kubernetes Manifest Deployment Engine
// @version 1.0
// @description save microservices and deploy services on kubernetes cluster
// @termsOfService http://swagger.io/terms/
// @contact.name Cloudplex Support
// @contact.url http://www.cloudplex.io/support
// @contact.email haseeb@cloudplex.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /ksd/api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name X-Auth-Token
func main() {

	e := gin.New()
	utils.InitFlags()
	constants.CacheObj = cache.New(5*time.Minute, 6*time.Minute)
	if constants.ServicePort == "" {
		constants.ServicePort = "8089"
	}
	c, _ := controllers.NewController()
	v1 := e.Group("/ksd/api/v1")
	{
		v1.Use(auth())
		v1.POST("/solution", c.DeploySolution)
		v1.GET("/solution", c.GetSolution)
		v1.GET("/solution/all", c.ListSolution)
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
		v1.GET("/kubehealthservice", c.GetKubernetesServiceHealth)

		v1.GET("/secret/:namespace/:name", c.GetRegistrySecret)
		v1.POST("/secret/:namespace/:name", c.CreateRegistrySecret)
		v1.DELETE("/secret/:namespace/:name", c.DeleteRegistrySecret)

		v1.GET("/getallnamespaces/:project_id/", c.GetAllNamespaces)
		v1.GET("/health", controllers.Health)
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	go grpcMain()
	e.Run(":" + constants.ServicePort)
	//e.Logger.Fatal(e.Start(":" + constants.ServicePort))

}

func grpcMain() {
	port := fmt.Sprintf(":%s", constants.ServiceGRPCPort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	svc := &core.Server{}
	pb.RegisterServiceServer(srv, svc)

	pb.RegisterK8SResourceServer(srv, svc)
	pb.RegisterAnnotateNamespaceServer(srv, svc)
	pb.RegisterPodKillServiceServer(srv, svc)
	//go handleclient()

	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleclient() {
	conn, err := grpc.Dial("localhost:8094", grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
	}

	_, err = pb.NewK8SResourceClient(conn).GetK8SResource(context.Background(), &pb.KubernetesResourceRequest{
		InfraId:   "11",
		CompanyId: "1111",
		Token:     "dfsdfsdf",
	})
	if err != nil {
		utils.Error.Println(err)
	}

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

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader(constants.AuthTokenKey)) == 0 {
			c.JSON(http.StatusUnauthorized, types.HTTPError{Message: "X-Auth-Token is required in Header"})
			c.Abort()
		}
		c.Next()
	}
}
