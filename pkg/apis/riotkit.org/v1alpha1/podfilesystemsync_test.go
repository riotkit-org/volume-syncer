package v1alpha1

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestPodFilesystemSync_IsPodMatching(t *testing.T) {
	// case A
	definitionA := PodFilesystemSync{}
	definitionA.Spec.PodSelector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"variant": "with-dynamic-directory-name",
		},
	}
	firstPod := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"riotkit.org/volume-syncing-operator": "true",
				"variant":                             "with-dynamic-directory-name",
			},
		},
	}

	// case B
	definitionB := PodFilesystemSync{}
	definitionB.Spec.PodSelector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"my-pod-label": "test",
		},
	}
	secondPod := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"riotkit.org/volume-syncing-operator": "true",
				"my-pod-label":                        "test",
			},
		},
	}

	assert.True(t, definitionA.IsPodMatching(&firstPod))
	assert.False(t, definitionA.IsPodMatching(&secondPod))

	assert.True(t, definitionB.IsPodMatching(&secondPod))
	assert.False(t, definitionB.IsPodMatching(&firstPod))
}
