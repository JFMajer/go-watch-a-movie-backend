package main

import (
	"backend/internal/models"
)

var movies = []models.Movie{
	{ID: 1, Title: "Highlander", Description: "An immortal Scottish swordsman must confront the last of his immortal opponent, a murderously brutal barbarian who lusts for the fabled 'Prize'.", Year: 1986, Runtime: 116, MPAARating: "R"},
	{ID: 2, Title: "The Shawshank Redemption", Description: "Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.", Year: 1994, Runtime: 142, MPAARating: "R"},
	{ID: 3, Title: "Inception", Description: "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.", Year: 2010, Runtime: 148, MPAARating: "PG-13"},
	{ID: 4, Title: "The Matrix", Description: "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.", Year: 1999, Runtime: 136, MPAARating: "R"},
	{ID: 5, Title: "Pulp Fiction", Description: "The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.", Year: 1994, Runtime: 154, MPAARating: "R"},
	// ... Add the remaining movies in a similar fashion
}
