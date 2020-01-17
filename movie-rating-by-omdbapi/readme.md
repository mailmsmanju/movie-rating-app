# Get Movie Rating App 
App is written in Golang, its command line App provides easy way to get movie ratings from OMDB web site by passing movie name as argument.

## Getting Started
To start using Get Movie Rating App , install 
* go 
* docker

## Build Steps

This project structure as follows 

        movie-rating-by-omdbapi
        +----Dockerfile
        +----GetMovieRating.go
        +----GetMovieRating_test.go
        +----ReadMe.md

## Building & Run docker:

 1) Export environemtn variable for local project path

        $export LOCAL_LIB_PATH = <project local path> // Ex. ${HOME}/movie-rating-by-omdbapi

 2) Create docker:

        $make build -f ${LOCAL_LIB_PATH}/bin/makefile -e "${LOCAL_LIB_PATH}"

 3) Run docker:

        $make run -f ${LOCAL_LIB_PATH}/makefile -e "title=<movie title>" // ex. star

## Compile from source code and Run test

   1) Build using go build command from project directory

            $ cd ${LOCAL_LIB_PATH}
            $ go build -o <nameofexecutable> main .
   2) Run 

            $ ./<nameofexecutable> -title=<title> // Ex. star  



Sample out put

  ------------------------------------ Your request ------------------------------------------

Get me Rotten Tomatoes rating for the movie ====>  star

 ------------------------- Please find your requested information ---------------------------

Movie Name is       ====> Star Wars: Episode IX - The Rise of Skywalker

Movie Rating Source ====> Rotten Tomatoes

Movie Rating        ====> 55%

-------------------------------------- Thank you --------------------------------------------  

        