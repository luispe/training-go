package main

import (
	"fmt"
	"sort"
)

type user struct {
	FirstName       string
	LastName        string
	Age             int
	FavoritesThings []string
}

type SortAge []user

func (sa SortAge) Len() int           { return len(sa) }
func (sa SortAge) Swap(i, j int)      { sa[i], sa[j] = sa[j], sa[i] }
func (sa SortAge) Less(i, j int) bool { return sa[i].Age < sa[j].Age }

type SortLastName []user

func (sln SortLastName) Len() int           { return len(sln) }
func (sln SortLastName) Swap(i, j int)      { sln[i], sln[j] = sln[j], sln[i] }
func (sln SortLastName) Less(i, j int) bool { return sln[i].LastName < sln[j].LastName }

func main() {
	u1 := user{
		FirstName: "Homer",
		LastName:  "Simpson",
		Age:       42,
		FavoritesThings: []string{
			"Beer",
			"Donuts",
			"Sofa",
		},
	}

	u2 := user{
		FirstName: "Barnie",
		LastName:  "Gomez",
		Age:       39,
		FavoritesThings: []string{
			"Beer",
			"Helicopter",
			"Bar",
		},
	}

	u3 := user{
		FirstName: "Ned",
		LastName:  "Flanders",
		Age:       40,
		FavoritesThings: []string{
			"Church",
			"Family",
			"Walk",
		},
	}

	users := []user{u1, u2, u3}
	for _, v := range users {
		fmt.Println(v.FirstName, v.LastName, v.Age)
		sort.Strings(v.FavoritesThings)
		for _, v := range v.FavoritesThings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("----------")
	fmt.Println("Sort users age")
	sort.Sort(SortAge(users))
	for _, v := range users {
		fmt.Println(v.FirstName, v.LastName, v.Age)
		for _, v := range v.FavoritesThings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("----------")
	fmt.Println("Sort users last name")
	sort.Sort(SortLastName(users))
	for _, v := range users {
		fmt.Println(v.FirstName, v.LastName, v.Age)
		for _, v := range v.FavoritesThings {
			fmt.Println("\t", v)
		}
	}
}
