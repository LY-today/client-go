package fake

import (
	"context"
	v1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNewSimpleClientset(t *testing.T) {
	client := NewSimpleClientset()
	client.CoreV1().Pods("default").Create(context.Background(), &v1.Pod{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "pod-1",
			Namespace: "default",
		},
	}, meta_v1.CreateOptions{})
	client.CoreV1().Pods("default").Create(context.Background(), &v1.Pod{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "pod-2",
			Namespace: "default",
		},
	}, meta_v1.CreateOptions{})
	_ = client.CoreV1().Pods("default").EvictV1(context.Background(), &policy.Eviction{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: "pod-2",
		},
	})
	pods, err := client.CoreV1().Pods("default").List(context.Background(), meta_v1.ListOptions{})
	//err: item[0]: can't assign or convert v1beta1.Eviction into v1.Pod
	if err != nil {
		t.Errorf("TestNewSimpleClientset() res = %v", err.Error())
	}else {
		t.Logf("TestNewSimpleClientset() res = %v", pods)
	}
}
