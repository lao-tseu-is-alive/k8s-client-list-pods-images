package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"os"
	"path/filepath"
	"strings"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

var (
	APP        = "k8sClientListPodsImages"
	AppSnake   = "k8s-client-list-pods-images"
	VERSION    = "0.0.1"
	REPOSITORY = "github.com/lao-tseu-is-alive/k8s-client-list-pods-images"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	nameSpace2Use := flag.String("n", "", "search pattern to select namespaces")
	verbose := flag.Bool("v", false, "be more verbose on the output")
	header := flag.Bool("header", false, "display the first columns header")
	flag.Parse()
	l := log.New(os.Stdout, fmt.Sprintf("%s ", APP), log.Ldate|log.Ltime|log.Lshortfile)
	if *verbose == true {
		l.Printf("INFO: 'Starting %s v:%s '", AppSnake, VERSION)
		l.Printf("INFO: 'Repository url: https://%s'", REPOSITORY)
	}

	namespace := *nameSpace2Use
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		l.Printf("ERROR: doing BuildConfigFromFlags trying to retrieve kubeconfig err:%v \n", err)
		panic(err.Error())
	}

	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		l.Printf("ERROR: doing kubernetes.NewForConfig trying to build client set from kubeconfig err:%v \n", err)
		panic(err.Error())
	}

	namespaces, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		l.Printf("ERROR: trying to get the list of namespaces err:%v \n", err)
		panic(err.Error())
	}
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	if *verbose == true {
		fmt.Printf("## Found %d pods in all the cluster\n", len(pods.Items))
		fmt.Printf("## parsing namespace list to find for : %s ##\n", namespace)
		fmt.Print("## NumPods \t Namespace  ##\n")
	}
	for _, n := range namespaces.Items {
		currentNameSpace := n.Name
		pods, err := clientSet.CoreV1().Pods(currentNameSpace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			l.Printf("ERROR: trying to get the number of pods in namespace %s err:%v \n", currentNameSpace, err)
			panic(err.Error())
		}
		if *verbose == true {
			fmt.Printf("## %d \t %s\n", len(pods.Items), n.Name)
		}
		if strings.Contains(currentNameSpace, *nameSpace2Use) {
			if *header == true {
				fmt.Print("## Namespace\tPod Name\tPod Image\n")
			}
			for _, p := range pods.Items {
				currentPodName := p.Name
				//fmt.Printf("\n## POD : %s.%s  ##\n", currentNameSpace, currentPodName)
				// Examples for error handling:
				// - Use helper functions like e.g. errors.IsNotFound()
				// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
				currentPod, err := clientSet.CoreV1().Pods(currentNameSpace).Get(context.TODO(), currentPodName, metav1.GetOptions{})
				if errors.IsNotFound(err) {
					l.Printf("Pod %s in namespace %s not found\n", currentPodName, namespace)
				} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
					l.Printf("Error getting pod %s in namespace %s: %v\n",
						currentPodName, namespace, statusError.ErrStatus.Message)
				} else if err != nil {
					panic(err.Error())
				} else {
					for _, c := range currentPod.Spec.Containers {
						fmt.Printf("%s\t%s\t%s\n", currentNameSpace, currentPodName, c.Image)
					}
				}

			}
		}
	}
}
