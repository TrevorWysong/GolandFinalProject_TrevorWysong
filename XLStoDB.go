package main

import (
	"database/sql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

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

func main() {
	openedExcel, err := excelize.OpenFile("games-features.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	column, err := openedExcel.GetRows("games-features")
	for _, column := range column {
		fmt.Println(column, "\t")
		A := column[0]
		B := column[1]
		C := column[2]
		D := column[3]
		E := column[4]
		F := column[5]
		G := column[6]
		H := column[7]
		I := column[8]
		J := column[9]
		K := column[10]
		L := column[11]
		M := column[12]
		N := column[13]
		O := column[14]
		P := column[15]
		Q := column[16]
		R := column[17]
		S := column[18]
		T := column[19]
		U := column[20]
		V := column[21]
		W := column[22]
		X := column[23]
		Y := column[24]
		Z := column[25]
		AA := column[26]
		AB := column[27]
		AC := column[28]
		AD := column[29]
		AE := column[30]
		AF := column[31]
		AG := column[32]
		AH := column[33]
		AI := column[34]
		AJ := column[35]
		AK := column[36]
		AL := column[37]
		AM := column[38]
		AN := column[39]
		AO := column[40]
		AP := column[41]
		AQ := column[42]
		AR := column[43]
		AS := column[44]
		AT := column[45]
		AU := column[46]
		AV := column[47]
		AW := column[48]
		AX := column[49]
		AY := column[50]
		AZ := column[51]
		BA := column[52]
		BB := column[53]
		BC := column[54]
		BD := column[55]
		BE := column[56]
		BF := column[57]
		BG := column[58]
		BH := column[59]
		BI := column[60]
		BJ := column[61]
		BK := column[62]
		BL := column[63]
		BM := column[64]
		BN := column[65]
		BO := column[66]
		BP := column[67]
		BQ := column[68]
		BR := column[69]
		BS := column[70]
		BT := column[71]
		BU := column[72]
		BV := column[73]
		BW := column[74]
		BX := column[75]
		BY := column[76]
		BZ := column[77]
		gameDatabase := OpenDataBase("./GameData.db")
		create_table(gameDatabase)
		defer gameDatabase.Close()
		insert_data(gameDatabase, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z,
			AA, AB, AC, AD, AE, AF, AG, AH, AI, AJ, AK, AL, AM, AN, AO, AP, AQ, AR, AS, AT, AU, AV, AW, AX, AY, AZ,
			BA, BB, BC, BD, BE, BF, BG, BH, BI, BJ, BK, BL, BM, BN, BO, BP, BQ, BR, BS, BT, BU, BV, BW, BX, BY, BZ)
	}
	gameDatabase := OpenDataBase("./GameData.db")
	defer gameDatabase.Close()
	create_table(gameDatabase)
}
