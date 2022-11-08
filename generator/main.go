package main

import (
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/printers"
)

func main() {
	for i := 0; i < 1000; i++ {
		name := fmt.Sprintf("my-data-%d", i)
		cm := configMap(name, "gamestore")

		fileName := name + ".yaml"
		serialize(&cm, fileName)
	}
}

func serialize(cm *corev1.ConfigMap, filename string) {
	newFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	y := printers.YAMLPrinter{}
	defer newFile.Close()
	y.PrintObj(cm, newFile)
}

func configMap(name, ns string) corev1.ConfigMap {
	cm := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Data: map[string]string{
			"fake-data":   "data-line-1",
			"fake-data-2": "data-line-2",
		},
	}

	return cm
}
