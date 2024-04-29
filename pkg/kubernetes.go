package pkg

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"

	v1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func NewKubernetesClient() (*kubernetes.Clientset, error) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return clientset, nil
}

func GetKubernetesNodes() {
	clientset, err := NewKubernetesClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for y, nodes := range nodes.Items {
		fmt.Printf("[%d] %s\n", y, nodes.GetName())
	}

}

func CreateDeployment() {
	clientset, err := NewKubernetesClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func CreateService() {
	clientset, err := NewKubernetesClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo",
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt(8080),
				},
			},
			Selector: map[string]string{
				"app": "demo",
			},
		},
	}

	// Create Service
	fmt.Println("Creating service...")
	result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service%q.\n", result.GetObjectMeta().GetName())
}

func CreateIngress() {
	clientset, err := NewKubernetesClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	ingressClient := clientset.NetworkingV1beta1().Ingresses("default")

	ingress := &v1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind: "Ingress",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ingress-rule-created-from-go",
			Namespace: "default",
		},
		Spec: v1beta1.IngressSpec{
			Backend: &v1beta1.IngressBackend{
				ServiceName: "demo",

				ServicePort: intstr.FromInt(80),
			},
			Rules: []v1beta1.IngressRule{{Host: "ex.Artfolio.com"}},
		},
	}
	fmt.Println("Creating ingress...")
	result, err := ingressClient.Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Created Ingress %q.\n", result.GetObjectMeta().GetName())
}

func int32Ptr(i int32) *int32 { return &i }
