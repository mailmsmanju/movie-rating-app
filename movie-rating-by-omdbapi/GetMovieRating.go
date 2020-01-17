/*
# GetMovieRating go file accepts movie name as argument from command line and returns Rotten Tomatoes rating for the movie on command line
# Author: Manju Subramanyam
# Date: Jan 15th 2020
# Copyright (c) 2018 Cisco Systems. All rights reserved.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// START
func main() {

	// STEP 1 - Reading input entered by user in termial and replacing spaces.
	newmovienameonly := spaceReplace(inputReqArg())

	fmt.Println("\n ------------------------------------ Your request ------------------------------------------")
	fmt.Println("\nGet me Rotten Tomatoes rating for the movie ====> ", newmovienameonly)

	//STEP 2 - Check argument value, if empty return appropriate message to user on terminal.
	if len(newmovienameonly) < 1 {
		fmt.Println("\n" + newmovienameonly + "movie name can not be empty, please enter correct movie name and try again...Thanks")
	}

	//STEP 3 - if not empty request omdb pai for movie ratings.
	response, err := http.Get("http://www.omdbapi.com/?t=" + newmovienameonly + "&y=2019&apikey=a236cc5b")

	//STEP 4 - if nil return error message back to user on terminal.
	if err != nil {
		fmt.Println("Get request failed  \n", err)

	} else {

		//STEP 5 - process the rest api response jason, serach and return Rotten Tomatoes rating for the movie to the user.
		data, _ := ioutil.ReadAll(response.Body)
		var movie Movie
		json.Unmarshal(data, &movie)

		for i := 0; i < len(movie.Ratings); i++ {

			if IsValidSource(movie.Ratings[i].Source) {
				fmt.Println("\n ------------------------- Please find your requested information ---------------------------")
				fmt.Println("\nMovie Name is       ====> " + movie.Title)
				fmt.Println("\nMovie Rating Source ====> " + movie.Ratings[i].Source)
				fmt.Println("\nMovie Rating        ====> " + movie.Ratings[i].Value)

			}
		}

	}

	fmt.Println("\n-------------------------------------- Thank you --------------------------------------------")

}

//inputReqArg reads movie name entered as argument by user.
func inputReqArg() string {
	var inputMovieName string
	if len(os.Args) > 1 {
		inputMovieName = os.Args[1]
	}
	return inputMovieName
}

//spaceReplace replaces emty spaces in string, Replace the " " with a "%20"
func spaceReplace(replaceString string) string {
	result := strings.Replace(replaceString, " ", "%20", -1)
	return result
}

//IsValidSource to check if source type is "Rotten Tomatoes" or not. takes string as input and returns boolean value true/false.
func IsValidSource(source string) bool {
	switch source {
	case
		"Rotten Tomatoes":
		return true
	}
	return false
}

//Movie struct is to save movie informaion from the response jason and parse it to get movie rating and source.
type Movie struct {
	Title      string    `json:"Title"`
	Year       string    `json:"Year"`
	Rated      string    `json:"Rated"`
	Released   string    `json:"Released"`
	Runtime    string    `json:"Runtime"`
	Genre      string    `json:"Genre"`
	Director   string    `json:"Director"`
	Writer     string    `json:"Writer"`
	Actors     string    `json:"Actors"`
	Plot       string    `json:"Plot"`
	Language   string    `json:"Language"`
	Country    string    `json:"Country"`
	Awards     string    `json:"Awards"`
	Poster     string    `json:"Poster"`
	Ratings    []Ratings `json:"Ratings"`
	Metascore  string    `json:"Metascore"`
	ImdbRating string    `json:"imdbRating"`
	ImdbVotes  string    `json:"imdbVotes"`
	ImdbID     string    `json:"imdbID"`
	Type       string    `json:"Type"`
	DVD        string    `json:"DVD"`
	BoxOffice  string    `json:"BoxOffice"`
	Production string    `json:"Production"`
	Website    string    `json:"Website"`
	Response   bool      `json:"Response"`
}

//Ratings struct is to save rating array informaion from the response jason and parse it to get movie rating and source.
type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
