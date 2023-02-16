# k8s-client-list-pods-images
k8s-client-list-pods-images is a simple [kubernetes](https://kubernetes.io/) client written in Go that list the pods names and container images in your cluster.

## usage :

    ./k8s-client-list-pods-images -n go-testing

    go-testing	go-cloud-k8s-info-595ff998d9-6hncq	ghcr.io/lao-tseu-is-alive/go-cloud-k8s-info:v0.4.10
    go-testing	go-cloud-k8s-info-595ff998d9-twr82	ghcr.io/lao-tseu-is-alive/go-cloud-k8s-info:v0.4.10
    go-testing	go-cloud-k8s-shell-5d789684b6-2qlq7	ghcr.io/lao-tseu-is-alive/go-cloud-k8s-shell:v0.1.13
    go-testing	go-cloud-k8s-shell-5d789684b6-9tq96	ghcr.io/lao-tseu-is-alive/go-cloud-k8s-shell:v0.1.13

will list the namespaces, pod names and pod images found in namespace containing go-testing,   every field separated by a tab.

with no namespaces search parameter it will list all pods, in all namespace

    ./k8s-client-list-pods-images -h
  
    Usage of ./k8s-client-list-pods-images:
    -header
        	display the first columns header
    -kubeconfig string
        	(optional) absolute path to the kubeconfig file (default "/home/cgil/.kube/config")
    -n string
        	search pattern to select namespaces
    -v	be more verbose on the output

## more info:
   
 + [Kubernetes](https://kubernetes.io/), also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications.
 + We use the official [Go client library for kubernetes]( https://github.com/kubernetes/client-go).
 + This code is to be used from outside the cluster with your corresponding kubeconfig