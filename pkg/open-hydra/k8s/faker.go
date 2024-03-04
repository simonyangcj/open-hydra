package k8s

import (
	"fmt"
	"open-hydra/pkg/open-hydra/apis"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Fake struct {
	namespacedPod     map[string][]coreV1.Pod
	namespacedDeploy  map[string][]appsV1.Deployment
	namespacedService map[string][]coreV1.Service
	labelPod          map[string][]coreV1.Pod
	labelDeploy       map[string][]appsV1.Deployment
	labelService      map[string][]coreV1.Service
}

func (f *Fake) Init() {
	f.namespacedPod = make(map[string][]coreV1.Pod)
	f.namespacedDeploy = make(map[string][]appsV1.Deployment)
	f.namespacedService = make(map[string][]coreV1.Service)
	f.labelPod = make(map[string][]coreV1.Pod)
	f.labelDeploy = make(map[string][]appsV1.Deployment)
	f.labelService = make(map[string][]coreV1.Service)
}

func (f *Fake) ListDeploymentWithLabel(label, namespace string, client *kubernetes.Clientset) ([]appsV1.Deployment, error) {
	var result []appsV1.Deployment
	if _, ok := f.labelDeploy[label]; ok {
		result = f.labelDeploy[label]
	}
	return result, nil
}
func (f *Fake) ListPodWithLabel(label, namespace string, client *kubernetes.Clientset) ([]coreV1.Pod, error) {
	var result []coreV1.Pod
	if _, ok := f.labelPod[label]; ok {
		result = f.labelPod[label]
	}
	return result, nil
}
func (f *Fake) ListPod(namespace string, client *kubernetes.Clientset) ([]coreV1.Pod, error) {
	var result []coreV1.Pod
	if _, ok := f.namespacedPod[namespace]; ok {
		result = f.namespacedPod[namespace]
	}
	return result, nil
}
func (f *Fake) GetUserPods(label, namespace string, client *kubernetes.Clientset) ([]coreV1.Pod, error) {
	var result []coreV1.Pod
	if _, ok := f.labelPod[label]; ok {
		result = f.labelPod[label]
	}
	return result, nil
}
func (f *Fake) ListDeployment(namespace string, client *kubernetes.Clientset) ([]appsV1.Deployment, error) {
	var result []appsV1.Deployment
	if _, ok := f.namespacedDeploy[namespace]; ok {
		result = f.namespacedDeploy[namespace]
	}
	return result, nil
}
func (f *Fake) ListService(namespace string, client *kubernetes.Clientset) ([]coreV1.Service, error) {
	var result []coreV1.Service
	if _, ok := f.namespacedService[namespace]; ok {
		result = f.namespacedService[namespace]
	}
	return result, nil
}
func (f *Fake) DeleteUserDeployment(label, namespace string, client *kubernetes.Clientset) error {
	delete(f.labelDeploy, label)
	return nil
}
func (f *Fake) CreateDeployment(cpu, memory, image, namespace, studentID, ideType string, volumes []apis.VolumeMount, gpuSet apis.GpuSet, client *kubernetes.Clientset) error {
	label := fmt.Sprintf("%s=%s", OpenHydraUserLabelKey, studentID)
	f.labelDeploy[label] = append(f.labelDeploy[label], appsV1.Deployment{})
	f.namespacedDeploy[namespace] = append(f.namespacedDeploy[namespace], appsV1.Deployment{})
	return nil
}
func (f *Fake) CreateService(namespace, studentID, ideType string, client *kubernetes.Clientset) error {
	label := fmt.Sprintf("%s=%s", OpenHydraUserLabelKey, studentID)
	f.labelService[label] = append(f.labelService[label], coreV1.Service{})
	f.namespacedService[namespace] = append(f.namespacedService[namespace], coreV1.Service{})
	return nil
}
func (f *Fake) DeleteUserService(label, namespace string, client *kubernetes.Clientset) error {
	delete(f.labelService, label)
	return nil
}
func (f *Fake) GetUserService(label, namespace string, client *kubernetes.Clientset) (*coreV1.Service, error) {
	var result *coreV1.Service
	if _, ok := f.labelService[label]; ok {
		result = &f.labelService[label][0]
	} else {
		return nil, fmt.Errorf("service not found")
	}
	return result, nil
}
func (f *Fake) DeleteUserReplicaSet(label, namespace string, client *kubernetes.Clientset) error {
	return nil
}
func (f *Fake) DeleteUserPod(label, namespace string, client *kubernetes.Clientset) error {
	delete(f.labelPod, label)
	return nil
}
