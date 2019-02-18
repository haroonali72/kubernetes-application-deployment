package kubernetes

import (
	"encoding/json"
	"k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetesTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type ConfigMap struct {
	kubeClient *kubernetes.Clientset
}

func NewConfigLauncher(c *kubernetes.Clientset) *ConfigMap {
	this := new(ConfigMap)
	this.kubeClient = c
	return this
}
func (cm *ConfigMap) LaunchSideCarConfigMap() {

}
func (cm *ConfigMap) CreateConfigMap(configMap v1.ConfigMap) (*v1.ConfigMap, error) {

	return cm.kubeClient.CoreV1().ConfigMaps(configMap.ObjectMeta.Namespace).Create(&configMap)
}
func (cm *ConfigMap) PatchConfigMap(configMap v1.ConfigMap) (*v1.ConfigMap, error) {
	r, err := json.Marshal(configMap)
	if err != nil {
		return nil, err
	}
	return cm.kubeClient.CoreV1().ConfigMaps(configMap.ObjectMeta.Namespace).Patch(configMap.Name, kubernetesTypes.StrategicMergePatchType, r)
}
func (cm *ConfigMap) UpdateConfigMap(configMap *v1.ConfigMap) (*v1.ConfigMap, error) {

	return cm.kubeClient.CoreV1().ConfigMaps(configMap.ObjectMeta.Namespace).Update(configMap)
}
func (cm *ConfigMap) DeleteConfigMap(name, namespace string) error {
	return cm.kubeClient.CoreV1().ConfigMaps(namespace).Delete(name, &v12.DeleteOptions{})
}
func (cm *ConfigMap) GetConfigMap(name, namespace string) (*v1.ConfigMap, error) {
	return cm.kubeClient.CoreV1().ConfigMaps(namespace).Get(name, v12.GetOptions{})
}
func (cm *ConfigMap) GetAllConfigMap(namespace string) (*v1.ConfigMapList, error) {
	return cm.kubeClient.CoreV1().ConfigMaps(namespace).List(v12.ListOptions{})
}
