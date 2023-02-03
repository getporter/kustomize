package kustomize

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getSecret(ctx context.Context, client kubernetes.Interface, namespace, name, key string) ([]byte, error) {
	if namespace == "" {
		namespace = "default"
	}
	secret, err := client.CoreV1().Secrets(namespace).Get(ctx, name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		return nil, fmt.Errorf("error getting secret %s from namespace %s: %s", name, namespace, err)
	}
	val, ok := secret.Data[key]
	if !ok {
		return nil, fmt.Errorf("couldn't find key %s in secret", key)
	}
	return val, nil
}
