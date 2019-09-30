package kubernetes

import (
	"encoding/json"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetesTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"time"
)

type ServicesLauncher struct {
	kubeClient *kubernetes.Clientset
}

func NewServicesLauncher(c *kubernetes.Clientset) *ServicesLauncher {
	this := new(ServicesLauncher)
	this.kubeClient = c
	return this
}
func (p *ServicesLauncher) LaunchService(req *v1.Service) (svc *v1.Service, err error) {
	if req.Namespace == "" {
		req.Namespace = "default"
	}
	for svc == nil && err == nil {
		time.Sleep(1 * time.Second)
		svc, err = p.kubeClient.CoreV1().Services(req.Namespace).Create(req)
	}
	return svc, err
}

func (p *ServicesLauncher) PatchService(req *v1.Service) (svc *v1.Service, err error) {
	r, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	for svc == nil && err == nil {
		time.Sleep(1 * time.Second)
		svc, err = p.kubeClient.CoreV1().Services(req.Namespace).Patch(req.Name, kubernetesTypes.StrategicMergePatchType, r)
	}
	return svc, err
}
func (p *ServicesLauncher) UpdateService(req *v1.Service) (*v1.Service, error) {
	return p.kubeClient.CoreV1().Services(req.Namespace).Update(req)
}
func (p *ServicesLauncher) DeleteServices(serviceName, namespace string) error {
	return p.kubeClient.CoreV1().Services(namespace).Delete(serviceName, &metav1.DeleteOptions{})
}
func (p *ServicesLauncher) GetService(name, namespace string) (svc *v1.Service, err error) {
	for svc == nil && err == nil {
		time.Sleep(1 * time.Second)
		svc, err = p.kubeClient.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	}
	return svc, err
}
func (p *ServicesLauncher) GetAllServices(namespace string) (*v1.ServiceList, error) {
	return p.kubeClient.CoreV1().Services(namespace).List(metav1.ListOptions{})
}
