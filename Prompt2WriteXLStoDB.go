package main

import (
	"database/sql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	openedExcel, err := excelize.OpenFile("games-features.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	rows, err := openedExcel.GetRows("games-features")
	for _, row := range rows {
		fmt.Println(row, "\t")
		A := row[0]
		B := row[1]
		C := row[2]
		D := row[3]
		E := row[4]
		F := row[5]
		G := row[6]
		H := row[7]
		I := row[8]
		J := row[9]
		K := row[10]
		L := row[11]
		M := row[12]
		N := row[13]
		O := row[14]
		P := row[15]
		Q := row[16]
		R := row[17]
		S := row[18]
		T := row[19]
		U := row[20]
		V := row[21]
		W := row[22]
		X := row[23]
		Y := row[24]
		Z := row[25]
		AA := row[26]
		AB := row[27]
		AC := row[28]
		AD := row[29]
		AE := row[30]
		AF := row[31]
		AG := row[32]
		AH := row[33]
		AI := row[34]
		AJ := row[35]
		AK := row[36]
		AL := row[37]
		AM := row[38]
		AN := row[39]
		AO := row[40]
		AP := row[41]
		AQ := row[42]
		AR := row[43]
		AS := row[44]
		AT := row[45]
		AU := row[46]
		AV := row[47]
		AW := row[48]
		AX := row[49]
		AY := row[50]
		AZ := row[51]
		BA := row[52]
		BB := row[53]
		BC := row[54]
		BD := row[55]
		BE := row[56]
		BF := row[57]
		BG := row[58]
		BH := row[59]
		BI := row[60]
		BJ := row[61]
		BK := row[62]
		BL := row[63]
		BM := row[64]
		BN := row[65]
		BO := row[66]
		BP := row[67]
		BQ := row[68]
		BR := row[69]
		BS := row[70]
		BT := row[71]
		BU := row[72]
		BV := row[73]
		BW := row[74]
		BX := row[75]
		BY := row[76]
		BZ := row[77]
		myDatabase := OpenDataBase("./GameData.db")
		create_table(myDatabase)
		defer myDatabase.Close()
		insert_data(myDatabase, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z,
			AA, AB, AC, AD, AE, AF, AG, AH, AI, AJ, AK, AL, AM, AN, AO, AP, AQ, AR, AS, AT, AU, AV, AW, AX, AY, AZ,
			BA, BB, BC, BD, BE, BF, BG, BH, BI, BJ, BK, BL, BM, BN, BO, BP, BQ, BR, BS, BT, BU, BV, BW, BX, BY, BZ)
	}
	GameDB()
}
func GameDB() {
	myDatabase := OpenDataBase("./GameData.db")
	defer myDatabase.Close()
	create_table(myDatabase)
}

func OpenDataBase(dbfile string) *sql.DB {
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func create_table(database *sql.DB) {
	createStatement1 := "CREATE TABLE IF NOT EXISTS GameTable( " +
		"QueryID INTEGER," +
		"ResponseID INTEGER," +
		"QueryName TEXT," +
		"ResponseName TEXT," +

		"ReleaseDate DATE," +
		"RequiredAge INTEGER," +
		"DemoCount INTEGER," +
		"DeveloperCount INTEGER," +
		"DLCCount INTEGER," +
		"Metacritic INTEGER," +
		"MovieCount INTEGER," +
		"PackageCount INTEGER," +
		"RecommendationCount INTEGER," +
		"PublisherCount INTEGER," +
		"ScreenshotCount INTEGER," +
		"SteamSpyOwners INTEGER," +
		"SteamSpyOwnersVariance INTEGER," +
		"SteamSpyPlayersEstimate INTEGER," +
		"SteamSpyPlayersVariance INTEGER," +
		"AchievementCount INTEGER," +
		"AchievementHighlightedCount INTEGER," +
		"ControllerSupport BOOLEAN," +
		"IsFree BOOLEAN," +
		"FreeVerAvail BOOLEAN," +
		"PurchaseAvail BOOLEAN," +
		"SubscriptionAvail BOOLEAN," +
		"PlatformWindows BOOLEAN," +
		"PlatformLinux BOOLEAN," +
		"PlatformMac BOOLEAN," +
		"PCReqsHaveMin BOOLEAN," +
		"PCReqsHaveRec BOOLEAN," +
		"LinuxReqsHaveMin BOOLEAN," +
		"LinuxReqsHaveRec BOOLEAN," +
		"MacReqsHaveMin BOOLEAN," +
		"MacReqsHaveRec BOOLEAN," +
		"CategorySinglePlayer BOOLEAN," +
		"CategoryMultiplayer BOOLEAN," +
		"CategoryCoop BOOLEAN," +
		"CategoryMMO BOOLEAN," +
		"CategoryInAppPurchase BOOLEAN," +
		"CategoryIncludeSrcSDK BOOLEAN," +
		"CategoryIncludeLevelEditor BOOLEAN," +
		"CategoryVRSupport BOOLEAN," +
		"GenreIsNonGame BOOLEAN," +
		"GenreIsIndie BOOLEAN," +
		"GenreIsAction BOOLEAN," +
		"GenreIsAdventure BOOLEAN," +
		"GenreIsCasual BOOLEAN," +
		"GenreIsStrategy BOOLEAN," +
		"GenreIsRPG BOOLEAN," +
		"GenreIsSimulation BOOLEAN," +
		"GenreIsEarlyAccess BOOLEAN," +
		"GenreIsFreeToPlay BOOLEAN," +
		"GenreIsSports BOOLEAN," +
		"GenreIsRacing BOOLEAN," +
		"GenreIsMassivelyMultiplayer BOOLEAN," +
		"PriceCurrency TEXT," +
		"PriceInitial FLOAT," +
		"PriceFinal FLOAT," +
		"SupportEmail TEXT," +
		"SupportURL TEXT," +
		"AboutText TEXT," +
		"Background TEXT," +
		"ShortDescrip TEXT," +
		"DetailedDescrip TEXT," +
		"DRMNotice TEXT," +
		"ExtUserAcctNotice TEXT," +
		"HeaderImage TEXT," +
		"LegalNotice TEXT," +
		"Reviews TEXT," +
		"SupportedLanguages TEXT," +
		"Website TEXT," +
		"PCMinReqsText TEXT," +
		"PCRecReqsText TEXT," +
		"LinuxMinReqsText TEXT," +
		"LinuxRecReqsText TEXT," +
		"MacMinReqsText TEXT," +
		"MacRecReqsText TEXT);"
	database.Exec(createStatement1)
}
func insert_data(database *sql.DB, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P,
	Q, R, S, T, U, V, W, X, Y, Z, AA, AB, AC, AD, AE, AF, AG, AH, AI, AJ, AK, AL,
	AM, AN, AO, AP, AQ, AR, AS, AT, AU, AV, AW, AX, AY, AZ, BA, BB, BC, BD, BE, BF, BG,
	BH, BI, BJ, BK, BL, BM, BN, BO, BP, BQ, BR, BS, BT, BU, BV, BW, BX, BY, BZ string) {

	insertStatement := "INSERT INTO GameTable(QueryID,ResponseID,QueryName,ResponseName,ReleaseDate,RequiredAge,DemoCount,DeveloperCount,DLCCount,Metacritic,MovieCount,PackageCount,RecommendationCount,PublisherCount,ScreenshotCount,SteamSpyOwners,SteamSpyOwnersVariance,SteamSpyPlayersEstimate,SteamSpyPlayersVariance,AchievementCount,AchievementHighlightedCount,ControllerSupport,IsFree,FreeVerAvail,PurchaseAvail,SubscriptionAvail,PlatformWindows,PlatformLinux,PlatformMac,PCReqsHaveMin,PCReqsHaveRec,LinuxReqsHaveMin,LinuxReqsHaveRec,MacReqsHaveMin,MacReqsHaveRec,CategorySinglePlayer,CategoryMultiplayer,CategoryCoop,CategoryMMO,CategoryInAppPurchase,CategoryIncludeSrcSDK,CategoryIncludeLevelEditor,CategoryVRSupport,GenreIsNonGame,GenreIsIndie,GenreIsAction,GenreIsAdventure,GenreIsCasual,GenreIsStrategy,GenreIsRPG,GenreIsSimulation,GenreIsEarlyAccess,GenreIsFreeToPlay,GenreIsSports,GenreIsRacing,GenreIsMassivelyMultiplayer,PriceCurrency,PriceInitial,PriceFinal,SupportEmail,SupportURL,AboutText,Background,ShortDescrip,DetailedDescrip,DRMNotice,ExtUserAcctNotice,HeaderImage,LegalNotice,Reviews,SupportedLanguages,Website,PCMinReqsText,PCRecReqsText,LinuxMinReqsText,LinuxRecReqsText,MacMinReqsText,MacRecReqsText) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	prepped_statement, err := database.Prepare(insertStatement)
	if err != nil {
		log.Fatal(err)
	}
	prepped_statement.Exec(A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z, AA, AB, AC, AD,
		AE, AF, AG, AH, AI, AJ, AK, AL, AM, AN, AO, AP, AQ, AR, AS, AT, AU, AV, AW, AX, AY, AZ, BA, BB,
		BC, BD, BE, BF, BG, BH, BI, BJ, BK, BL, BM, BN, BO, BP, BQ, BR, BS, BT, BU, BV, BW, BX, BY, BZ)
}
