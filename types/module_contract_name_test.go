package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestModuleContractName_Match(t *testing.T) {
	tests := []struct {
		name      string
		mustMatch ModuleContractName
		candidate ModuleContractName
	}{
		{
			name: "match capability",
			mustMatch: ModuleContractName{
				Category: "capability",
				Provider: "aws",
				Platform: "*",
			},
			candidate: ModuleContractName{
				Category:    "capability",
				Subcategory: "ingress",
				Provider:    "aws",
				Platform:    "alb",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.mustMatch.Match(test.candidate)
			assert.True(t, result)
		})
	}
}

func TestCompareModuleContractName(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		wantLess bool
	}{
		{
			a:        "app:server/aws/ec2:beanstalk",
			b:        "app:server/aws/ec2:*",
			wantLess: true,
		},
		{
			a:        "app:container/aws/ecs:fargate",
			b:        "app:container/aws/ecs:*",
			wantLess: true,
		},
		{
			a:        "app:container/aws/ecs:fargate",
			b:        "app:container/aws/ecs",
			wantLess: true,
		},
		{
			a:        "app:container/aws/ecs:fargate",
			b:        "app:container/aws/ecs:*",
			wantLess: true,
		},
		{
			a:        "app:container/aws/ecs:fargate",
			b:        "app:container/aws/ecs:ec2",
			wantLess: false,
		},
		{
			a:        "app:container/aws/ec2",
			b:        "app:server/aws/ecs",
			wantLess: true,
		},
		{
			a:        "app:container/aws/k8s",
			b:        "app:container/gcp/k8s",
			wantLess: true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a, err := ParseModuleContractName(test.a)
			if err != nil {
				t.Fatalf("bad contract name (a): %s", err)
			}
			b, err := ParseModuleContractName(test.b)
			if err != nil {
				t.Fatalf("bad contract name (b): %s", err)
			}
			got := CompareModuleContractName(a, b)
			assert.Equal(t, test.wantLess, got)
		})
	}
}

func TestCompareModuleContractName_Sort(t *testing.T) {
	all := []string{
		"app:serverless/aws/lambda:*",
		"app:server/aws/ec2:*",
		"app:container/aws/ecs:*",
		"app:container/gcp/k8s:gke",
		"app:static-site/aws/s3:*",
		"app:server/aws/ec2:beanstalk",
	}
	want := []string{
		"app:container/aws/ecs:*",
		"app:container/gcp/k8s:gke",
		"app:server/aws/ec2:beanstalk",
		"app:server/aws/ec2:*",
		"app:serverless/aws/lambda:*",
		"app:static-site/aws/s3:*",
	}
	got := make([]string, len(all))
	copy(got, all)

	sort.SliceStable(got, func(i, j int) bool {
		a, _ := ParseModuleContractName(got[i])
		b, _ := ParseModuleContractName(got[j])
		return CompareModuleContractName(a, b)
	})

	assert.Equal(t, want, got)
}
