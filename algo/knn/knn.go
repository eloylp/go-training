package knn

import (
	"math"
)

type User struct {
	Name string
	Sports,
	Comedy,
	Terror,
	ScienceFiction,
	Documental float64
	Movies []string
}

func RecommendToFrom(user *User, people []*User) []string {
	var recommendations []string
	nearestDistance := math.Inf(1)
	for _, u := range people {
		if user.Name == u.Name {
			continue
		}
		d := Distance(user, u)
		if d < nearestDistance {
			recommendations = u.Movies
			nearestDistance = d
		}
	}
	return recommendations
}

func Distance(us1 *User, us2 *User) float64 {
	sports := math.Pow(us1.Sports-us2.Sports, 2)
	comedy := math.Pow(us1.Comedy-us2.Comedy, 2)
	terror := math.Pow(us1.Terror-us2.Terror, 2)
	science := math.Pow(us1.ScienceFiction-us2.ScienceFiction, 2)
	documental := math.Pow(us1.Documental-us2.Documental, 2)
	return math.Sqrt(sports + comedy + terror + science + documental)
}
