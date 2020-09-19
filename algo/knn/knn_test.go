package knn_test

import (
	"github.com/bmizerany/assert"
	"github.com/eloylp/go-training/algo/knn"
	"testing"
)

func TestRecommendToFrom(t *testing.T) {
	alice, bob, cristine, delia := people()
	assert.Equal(t, []string{"A day at paradise"}, knn.RecommendToFrom(alice, []*knn.User{alice, bob, cristine, delia}))
	assert.Equal(t, []string{"Mr bean"}, knn.RecommendToFrom(bob, []*knn.User{alice, bob, cristine, delia}))
	assert.Equal(t, []string{"The mandalorian"}, knn.RecommendToFrom(cristine, []*knn.User{alice, bob, cristine, delia}))
	assert.Equal(t, []string{"Clone wars", "How to train my dog"}, knn.RecommendToFrom(delia, []*knn.User{alice, bob, cristine, delia}))
}

func people() (*knn.User, *knn.User, *knn.User, *knn.User) {
	return &knn.User{
			Name:           "alice",
			Sports:         1,
			Comedy:         5,
			Terror:         1,
			ScienceFiction: 1,
			Documental:     1,
			Movies:         []string{"Mr bean"},
		},
		&knn.User{
			Name:           "bob",
			Sports:         1,
			Comedy:         5,
			Terror:         1,
			ScienceFiction: 1,
			Documental:     1,
			Movies:         []string{"A day at paradise"},
		},
		&knn.User{
			Name:           "cristine",
			Sports:         1,
			Comedy:         1,
			Terror:         1,
			ScienceFiction: 5,
			Documental:     5,
			Movies:         []string{"Clone wars", "How to train my dog"},
		},
		&knn.User{
			Name:           "delia",
			Sports:         1,
			Comedy:         1,
			Terror:         1,
			ScienceFiction: 5,
			Documental:     5,
			Movies:         []string{"The mandalorian"},
		}
}

func TestDistance(t *testing.T) {
	us1 := &knn.User{
		Name:           "us1",
		Sports:         4.5,
		Comedy:         2,
		Terror:         2,
		ScienceFiction: 1,
		Documental:     1,
	}
	us2 := &knn.User{
		Name:           "us1",
		Sports:         2.1,
		Comedy:         4,
		Terror:         4.7,
		ScienceFiction: 3,
		Documental:     1,
	}
	assert.Equal(t, 4.588027898781785, knn.Distance(us1, us2))
}
