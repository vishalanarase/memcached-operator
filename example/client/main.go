package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	clientset "github.com/vishalanarase/memcached-operator/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Error building kubeconfig:", err.Error())
		panic(err.Error())
	}

	client, err := clientset.NewForConfig(config)
	if err != nil {
		fmt.Println("Error creating clientset:", err.Error())
		panic(err.Error())
	}

	list, err := client.CacheV1().Memcacheds("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error listing memcacheds:", err.Error())
		panic(err.Error())
	}

	fmt.Printf("Found %d memcacheds in the cluster\n", len(list.Items))

	fmt.Println("Memcacheds:")
	for _, item := range list.Items {
		fmt.Printf("Name: %s, Size: %d\n", item.Name, item.Spec.Size)
	}

	fmt.Println("Done listing memcacheds")
	fmt.Println("Exiting...")
}
