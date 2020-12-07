package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
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

func getDefaultScreen(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an API. Search for game with endpoint /game/{gameName}")
}

func getGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	chosenGame := Game{}

	db, err := sql.Open("sqlite3", "./GameData.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM GameTable")
	if err != nil {
		panic(err)
	}

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

		if chosenGame.QueryName == params["gameName"] {
			json.NewEncoder(w).Encode(chosenGame)
			rows.Close()
			db.Close()
			return
		}
	}
	rows.Close()
	db.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func printGameInfo() {
	// Creates info section and prints the request for user input
	fmt.Println("...........................................................................")
	fmt.Println("This is a program to search for a game and receive JSON as a response. ")
	fmt.Println("...........................................................................")
	fmt.Println("http://localhost:3000/")
	fmt.Println("...........................................................................")
}

func main() {
	printGameInfo()

	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/", getDefaultScreen).Methods("GET")

	r.HandleFunc("/game/{gameName}", getGame).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}
