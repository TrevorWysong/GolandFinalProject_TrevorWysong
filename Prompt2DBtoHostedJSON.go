package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	"os"
	"strings"
)

// Student declares `Student` structure
type Game struct {
	QueryID                     int
	ResponseID                  int
	QueryName                   string
	ResponseName                string
	ReleaseDate                 string
	RequiredAge                 int
	DemoCount                   int
	DeveloperCount              int
	DLCCount                    int
	Metacritic                  int
	MovieCount                  int
	PackageCount                int
	RecommendationCount         int
	PublisherCount              int
	ScreenshotCount             int
	SteamSpyOwners              int
	SteamSpyOwnersVariance      int
	SteamSpyPlayersEstimate     int
	SteamSpyPlayersVariance     int
	AchievementCount            int
	AchievementHighlightedCount int
	ControllerSupport           bool
	IsFree                      bool
	FreeVerAvail                bool
	PurchaseAvail               bool
	SubscriptionAvail           bool
	PlatformWindows             bool
	PlatformLinux               bool
	PlatformMac                 bool
	PCReqsHaveMin               bool
	PCReqsHaveRec               bool
	LinuxReqsHaveMin            bool
	LinuxReqsHaveRec            bool
	MacReqsHaveMin              bool
	MacReqsHaveRec              bool
	CategorySinglePlayer        bool
	CategoryMultiplayer         bool
	CategoryCoop                bool
	CategoryMMO                 bool
	CategoryInAppPurchase       bool
	CategoryIncludeSrcSDK       bool
	CategoryIncludeLevelEditor  bool
	CategoryVRSupport           bool
	GenreIsNonGame              bool
	GenreIsIndie                bool
	GenreIsAction               bool
	GenreIsAdventure            bool
	GenreIsCasual               bool
	GenreIsStrategy             bool
	GenreIsRPG                  bool
	GenreIsSimulation           bool
	GenreIsEarlyAccess          bool
	GenreIsFreeToPlay           bool
	GenreIsSports               bool
	GenreIsRacing               bool
	GenreIsMassivelyMultiplayer bool
	PriceCurrency               string
	PriceInitial                float64
	PriceFinal                  float64
	SupportEmail                string
	SupportURL                  string
	AboutText                   string
	Background                  string
	ShortDescrip                string
	DetailedDescrip             string
	DRMNotice                   string
	ExtUserAcctNotice           string
	HeaderImage                 string
	LegalNotice                 string
	Reviews                     string
	SupportedLanguages          string
	Website                     string
	PCMinReqsText               string
	PCRecReqsText               string
	LinuxMinReqsText            string
	LinuxRecReqsText            string
	MacMinReqsText              string
	MacRecReqsText              string
}

func processDBtoStruct(gameName string) (Game, bool) {
	db, err := sql.Open("sqlite3", "./GameData.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM GameTable")
	if err != nil {
		panic(err)
	}

	var gameFound bool
	chosenGame := Game{}

	for rows.Next() {
		err = rows.Scan(&chosenGame.QueryID, &chosenGame.ResponseID, &chosenGame.QueryName, &chosenGame.ResponseName,
			&chosenGame.ReleaseDate, &chosenGame.RequiredAge, &chosenGame.DemoCount, &chosenGame.DeveloperCount,
			&chosenGame.DLCCount, &chosenGame.Metacritic, &chosenGame.MovieCount, &chosenGame.PackageCount,
			&chosenGame.RecommendationCount, &chosenGame.PublisherCount, &chosenGame.ScreenshotCount,
			&chosenGame.SteamSpyOwners, &chosenGame.SteamSpyOwnersVariance, &chosenGame.SteamSpyPlayersEstimate,
			&chosenGame.SteamSpyPlayersVariance, &chosenGame.AchievementCount, &chosenGame.AchievementHighlightedCount,
			&chosenGame.ControllerSupport, &chosenGame.IsFree, &chosenGame.FreeVerAvail, &chosenGame.PurchaseAvail,
			&chosenGame.SubscriptionAvail, &chosenGame.PlatformWindows, &chosenGame.PlatformLinux,
			&chosenGame.PlatformMac, &chosenGame.PCReqsHaveMin, &chosenGame.PCReqsHaveRec, &chosenGame.LinuxReqsHaveMin,
			&chosenGame.LinuxReqsHaveRec, &chosenGame.MacReqsHaveMin, &chosenGame.MacReqsHaveRec,
			&chosenGame.CategorySinglePlayer, &chosenGame.CategoryMultiplayer, &chosenGame.CategoryCoop,
			&chosenGame.CategoryMMO, &chosenGame.CategoryInAppPurchase, &chosenGame.CategoryIncludeSrcSDK,
			&chosenGame.CategoryIncludeLevelEditor, &chosenGame.CategoryVRSupport, &chosenGame.GenreIsNonGame,
			&chosenGame.GenreIsIndie, &chosenGame.GenreIsAction, &chosenGame.GenreIsAdventure, &chosenGame.GenreIsCasual,
			&chosenGame.GenreIsStrategy, &chosenGame.GenreIsRPG, &chosenGame.GenreIsSimulation,
			&chosenGame.GenreIsEarlyAccess, &chosenGame.GenreIsFreeToPlay, &chosenGame.GenreIsSports,
			&chosenGame.GenreIsRacing, &chosenGame.GenreIsMassivelyMultiplayer, &chosenGame.PriceCurrency,
			&chosenGame.PriceInitial, &chosenGame.PriceFinal, &chosenGame.SupportEmail, &chosenGame.SupportURL,
			&chosenGame.AboutText, &chosenGame.Background, &chosenGame.ShortDescrip, &chosenGame.DetailedDescrip,
			&chosenGame.DRMNotice, &chosenGame.ExtUserAcctNotice, &chosenGame.HeaderImage, &chosenGame.LegalNotice,
			&chosenGame.Reviews, &chosenGame.SupportedLanguages, &chosenGame.Website, &chosenGame.PCMinReqsText,
			&chosenGame.PCRecReqsText, &chosenGame.LinuxMinReqsText, &chosenGame.LinuxRecReqsText,
			&chosenGame.MacMinReqsText, &chosenGame.MacRecReqsText)

		if chosenGame.QueryName == gameName {
			gameFound = true
			rows.Close()
			db.Close()
			return chosenGame, gameFound
		}
	}
	rows.Close()
	db.Close()

	if err != nil {
		log.Fatal(err)
	}
	return chosenGame, gameFound
}

func getUserInput() string {
	// Creates info section and prints the request for user input
	fmt.Println("...........................................................................")
	fmt.Println("This is a program to search for a game and receive JSON as a response. ")
	fmt.Println("...........................................................................")
	fmt.Print("Enter the name of the song here: ")

	// Reads in user input
	// Used bufio because scanner wouldn't read in anything after a space(s)
	inputReader := bufio.NewReader(os.Stdin)
	nameOfGame, _ := inputReader.ReadString('\n')
	return strings.TrimSuffix(nameOfGame, "\n")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	nameOfGame := getUserInput()
	chosenGame, foundInputMatch := processDBtoStruct(nameOfGame)

	if foundInputMatch == true {
		gameJSON, _ := json.Marshal(chosenGame)
		fmt.Println(string(gameJSON))
		handleRequests()
	} else {
		fmt.Println("Game not in database.")
	}
}
