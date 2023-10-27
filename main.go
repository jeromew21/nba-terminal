package main

import (
	"encoding/json"
	"fmt"
	"io"
	//"log"
	"strings"
	//"net"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"net/http"
	//"reflect"
	//"os"
	"strconv"
	"time"
)

type DayQuery struct {
	Resource   string `json:"resource"`
	Parameters struct {
		LeagueID string `json:"LeagueID"`
		Season   string `json:"Season"`
		RegionID int    `json:"RegionID"`
		Date     string `json:"Date"`
		EST      string `json:"EST"`
	} `json:"parameters"`
	ResultSets []struct {
		NextGameList []struct {
			GameID         string `json:"gameID"`
			VtCity         string `json:"vtCity"`
			VtNickName     string `json:"vtNickName"`
			VtShortName    string `json:"vtShortName"`
			VtAbbreviation string `json:"vtAbbreviation"`
			HtCity         string `json:"htCity"`
			HtNickName     string `json:"htNickName"`
			HtShortName    string `json:"htShortName"`
			HtAbbreviation string `json:"htAbbreviation"`
			Date           string `json:"date"`
			Time           string `json:"time"`
			Day            string `json:"day"`
			Broadcasters   []struct {
				BroadcastID       string `json:"broadcastID"`
				BroadcasterName   string `json:"broadcasterName"`
				TapeDelayComments string `json:"tapeDelayComments"`
			} `json:"broadcasters"`
		} `json:"NextGameList,omitempty"`
		CompleteGameList []struct {
			GameID            string `json:"gameID"`
			VtCity            string `json:"vtCity"`
			VtNickName        string `json:"vtNickName"`
			VtShortName       string `json:"vtShortName"`
			VtAbbreviation    string `json:"vtAbbreviation"`
			HtCity            string `json:"htCity"`
			HtNickName        string `json:"htNickName"`
			HtShortName       string `json:"htShortName"`
			HtAbbreviation    string `json:"htAbbreviation"`
			Date              string `json:"date"`
			Time              string `json:"time"`
			Day               string `json:"day"`
			BroadcastID       string `json:"broadcastID"`
			BroadcasterName   string `json:"broadcasterName"`
			TapeDelayComments string `json:"tapeDelayComments"`
		} `json:"CompleteGameList,omitempty"`
	} `json:"resultSets"`
}

type PlayByPlayQuery struct {
	Meta struct {
		Version int    `json:"version"`
		Code    int    `json:"code"`
		Request string `json:"request"`
		Time    string `json:"time"`
	} `json:"meta"`
	Game struct {
		GameID  string `json:"gameId"`
		Actions []struct {
			ActionNumber             int       `json:"actionNumber"`
			Clock                    string    `json:"clock"`
			TimeActual               time.Time `json:"timeActual"`
			Period                   int       `json:"period"`
			PeriodType               string    `json:"periodType"`
			ActionType               string    `json:"actionType"`
			SubType                  string    `json:"subType"`
			Qualifiers               []any     `json:"qualifiers"`
			PersonID                 int       `json:"personId"`
			X                        any       `json:"x"`
			Y                        any       `json:"y"`
			Possession               int       `json:"possession"`
			ScoreHome                string    `json:"scoreHome"`
			ScoreAway                string    `json:"scoreAway"`
			Edited                   time.Time `json:"edited"`
			OrderNumber              int       `json:"orderNumber"`
			IsTargetScoreLastPeriod  bool      `json:"isTargetScoreLastPeriod"`
			XLegacy                  any       `json:"xLegacy"`
			YLegacy                  any       `json:"yLegacy"`
			IsFieldGoal              int       `json:"isFieldGoal"`
			Side                     any       `json:"side"`
			Description              string    `json:"description"`
			PersonIdsFilter          []any     `json:"personIdsFilter"`
			TeamID                   int       `json:"teamId,omitempty"`
			TeamTricode              string    `json:"teamTricode,omitempty"`
			Descriptor               string    `json:"descriptor,omitempty"`
			JumpBallRecoveredName    string    `json:"jumpBallRecoveredName,omitempty"`
			JumpBallRecoverdPersonID int       `json:"jumpBallRecoverdPersonId,omitempty"`
			PlayerName               string    `json:"playerName,omitempty"`
			PlayerNameI              string    `json:"playerNameI,omitempty"`
			JumpBallWonPlayerName    string    `json:"jumpBallWonPlayerName,omitempty"`
			JumpBallWonPersonID      int       `json:"jumpBallWonPersonId,omitempty"`
			JumpBallLostPlayerName   string    `json:"jumpBallLostPlayerName,omitempty"`
			JumpBallLostPersonID     int       `json:"jumpBallLostPersonId,omitempty"`
			Area                     string    `json:"area,omitempty"`
			AreaDetail               string    `json:"areaDetail,omitempty"`
			ShotDistance             float64   `json:"shotDistance,omitempty"`
			ShotResult               string    `json:"shotResult,omitempty"`
			ShotActionNumber         int       `json:"shotActionNumber,omitempty"`
			ReboundTotal             int       `json:"reboundTotal,omitempty"`
			ReboundDefensiveTotal    int       `json:"reboundDefensiveTotal,omitempty"`
			ReboundOffensiveTotal    int       `json:"reboundOffensiveTotal,omitempty"`
			PointsTotal              int       `json:"pointsTotal,omitempty"`
			AssistPlayerNameInitial  string    `json:"assistPlayerNameInitial,omitempty"`
			AssistPersonID           int       `json:"assistPersonId,omitempty"`
			AssistTotal              int       `json:"assistTotal,omitempty"`
			BlockPlayerName          string    `json:"blockPlayerName,omitempty"`
			BlockPersonID            int       `json:"blockPersonId,omitempty"`
			OfficialID               int       `json:"officialId,omitempty"`
			FoulPersonalTotal        int       `json:"foulPersonalTotal,omitempty"`
			FoulTechnicalTotal       int       `json:"foulTechnicalTotal,omitempty"`
			FoulDrawnPlayerName      string    `json:"foulDrawnPlayerName,omitempty"`
			FoulDrawnPersonID        int       `json:"foulDrawnPersonId,omitempty"`
			TurnoverTotal            int       `json:"turnoverTotal,omitempty"`
			StealPlayerName          string    `json:"stealPlayerName,omitempty"`
			StealPersonID            int       `json:"stealPersonId,omitempty"`
		} `json:"actions"`
	} `json:"game"`
}

type BoxScoreTeam struct {
	TeamID            int    `json:"teamId"`
	TeamName          string `json:"teamName"`
	TeamCity          string `json:"teamCity"`
	TeamTricode       string `json:"teamTricode"`
	Score             int    `json:"score"`
	InBonus           string `json:"inBonus"`
	TimeoutsRemaining int    `json:"timeoutsRemaining"`
	Periods           []struct {
		Period     int    `json:"period"`
		PeriodType string `json:"periodType"`
		Score      int    `json:"score"`
	} `json:"periods"`
	Players []struct {
		Status     string `json:"status"`
		Order      int    `json:"order"`
		PersonID   int    `json:"personId"`
		JerseyNum  string `json:"jerseyNum"`
		Position   string `json:"position,omitempty"`
		Starter    string `json:"starter"`
		Oncourt    string `json:"oncourt"`
		Played     string `json:"played"`
		Statistics struct {
			Assists                 int     `json:"assists"`
			Blocks                  int     `json:"blocks"`
			BlocksReceived          int     `json:"blocksReceived"`
			FieldGoalsAttempted     int     `json:"fieldGoalsAttempted"`
			FieldGoalsMade          int     `json:"fieldGoalsMade"`
			FieldGoalsPercentage    float64 `json:"fieldGoalsPercentage"`
			FoulsOffensive          int     `json:"foulsOffensive"`
			FoulsDrawn              int     `json:"foulsDrawn"`
			FoulsPersonal           int     `json:"foulsPersonal"`
			FoulsTechnical          int     `json:"foulsTechnical"`
			FreeThrowsAttempted     int     `json:"freeThrowsAttempted"`
			FreeThrowsMade          int     `json:"freeThrowsMade"`
			FreeThrowsPercentage    float64 `json:"freeThrowsPercentage"`
			Minus                   float64 `json:"minus"`
			Minutes                 string  `json:"minutes"`
			MinutesCalculated       string  `json:"minutesCalculated"`
			Plus                    float64 `json:"plus"`
			PlusMinusPoints         float64 `json:"plusMinusPoints"`
			Points                  int     `json:"points"`
			PointsFastBreak         int     `json:"pointsFastBreak"`
			PointsInThePaint        int     `json:"pointsInThePaint"`
			PointsSecondChance      int     `json:"pointsSecondChance"`
			ReboundsDefensive       int     `json:"reboundsDefensive"`
			ReboundsOffensive       int     `json:"reboundsOffensive"`
			ReboundsTotal           int     `json:"reboundsTotal"`
			Steals                  int     `json:"steals"`
			ThreePointersAttempted  int     `json:"threePointersAttempted"`
			ThreePointersMade       int     `json:"threePointersMade"`
			ThreePointersPercentage float64 `json:"threePointersPercentage"`
			Turnovers               int     `json:"turnovers"`
			TwoPointersAttempted    int     `json:"twoPointersAttempted"`
			TwoPointersMade         int     `json:"twoPointersMade"`
			TwoPointersPercentage   float64 `json:"twoPointersPercentage"`
		} `json:"statistics"`
		Name                  string `json:"name"`
		NameI                 string `json:"nameI"`
		FirstName             string `json:"firstName"`
		FamilyName            string `json:"familyName"`
		NotPlayingReason      string `json:"notPlayingReason,omitempty"`
		NotPlayingDescription string `json:"notPlayingDescription,omitempty"`
	} `json:"players"`
	Statistics struct {
		Assists                      int     `json:"assists"`
		AssistsTurnoverRatio         float64 `json:"assistsTurnoverRatio"`
		BenchPoints                  int     `json:"benchPoints"`
		BiggestLead                  int     `json:"biggestLead"`
		BiggestScoringRun            int     `json:"biggestScoringRun"`
		BiggestScoringRunScore       string  `json:"biggestScoringRunScore"`
		Blocks                       int     `json:"blocks"`
		BlocksReceived               int     `json:"blocksReceived"`
		FastBreakPointsAttempted     int     `json:"fastBreakPointsAttempted"`
		FastBreakPointsMade          int     `json:"fastBreakPointsMade"`
		FastBreakPointsPercentage    float64 `json:"fastBreakPointsPercentage"`
		FieldGoalsAttempted          int     `json:"fieldGoalsAttempted"`
		FieldGoalsEffectiveAdjusted  float64 `json:"fieldGoalsEffectiveAdjusted"`
		FieldGoalsMade               int     `json:"fieldGoalsMade"`
		FieldGoalsPercentage         float64 `json:"fieldGoalsPercentage"`
		FoulsOffensive               int     `json:"foulsOffensive"`
		FoulsDrawn                   int     `json:"foulsDrawn"`
		FoulsPersonal                int     `json:"foulsPersonal"`
		FoulsTeam                    int     `json:"foulsTeam"`
		FoulsTechnical               int     `json:"foulsTechnical"`
		FoulsTeamTechnical           int     `json:"foulsTeamTechnical"`
		FreeThrowsAttempted          int     `json:"freeThrowsAttempted"`
		FreeThrowsMade               int     `json:"freeThrowsMade"`
		FreeThrowsPercentage         float64 `json:"freeThrowsPercentage"`
		LeadChanges                  int     `json:"leadChanges"`
		Minutes                      string  `json:"minutes"`
		MinutesCalculated            string  `json:"minutesCalculated"`
		Points                       int     `json:"points"`
		PointsAgainst                int     `json:"pointsAgainst"`
		PointsFastBreak              int     `json:"pointsFastBreak"`
		PointsFromTurnovers          int     `json:"pointsFromTurnovers"`
		PointsInThePaint             int     `json:"pointsInThePaint"`
		PointsInThePaintAttempted    int     `json:"pointsInThePaintAttempted"`
		PointsInThePaintMade         int     `json:"pointsInThePaintMade"`
		PointsInThePaintPercentage   float64 `json:"pointsInThePaintPercentage"`
		PointsSecondChance           int     `json:"pointsSecondChance"`
		ReboundsDefensive            int     `json:"reboundsDefensive"`
		ReboundsOffensive            int     `json:"reboundsOffensive"`
		ReboundsPersonal             int     `json:"reboundsPersonal"`
		ReboundsTeam                 int     `json:"reboundsTeam"`
		ReboundsTeamDefensive        int     `json:"reboundsTeamDefensive"`
		ReboundsTeamOffensive        int     `json:"reboundsTeamOffensive"`
		ReboundsTotal                int     `json:"reboundsTotal"`
		SecondChancePointsAttempted  int     `json:"secondChancePointsAttempted"`
		SecondChancePointsMade       int     `json:"secondChancePointsMade"`
		SecondChancePointsPercentage float64 `json:"secondChancePointsPercentage"`
		Steals                       int     `json:"steals"`
		ThreePointersAttempted       int     `json:"threePointersAttempted"`
		ThreePointersMade            int     `json:"threePointersMade"`
		ThreePointersPercentage      float64 `json:"threePointersPercentage"`
		TimeLeading                  string  `json:"timeLeading"`
		TimesTied                    int     `json:"timesTied"`
		TrueShootingAttempts         float64 `json:"trueShootingAttempts"`
		TrueShootingPercentage       float64 `json:"trueShootingPercentage"`
		Turnovers                    int     `json:"turnovers"`
		TurnoversTeam                int     `json:"turnoversTeam"`
		TurnoversTotal               int     `json:"turnoversTotal"`
		TwoPointersAttempted         int     `json:"twoPointersAttempted"`
		TwoPointersMade              int     `json:"twoPointersMade"`
		TwoPointersPercentage        float64 `json:"twoPointersPercentage"`
	} `json:"statistics"`
}

type BoxScoreQuery struct {
	Meta struct {
		Version int    `json:"version"`
		Code    int    `json:"code"`
		Request string `json:"request"`
		Time    string `json:"time"`
	} `json:"meta"`
	Game struct {
		GameID            string    `json:"gameId"`
		GameTimeLocal     string    `json:"gameTimeLocal"`
		GameTimeUTC       time.Time `json:"gameTimeUTC"`
		GameTimeHome      string    `json:"gameTimeHome"`
		GameTimeAway      string    `json:"gameTimeAway"`
		GameEt            string    `json:"gameEt"`
		Duration          int       `json:"duration"`
		GameCode          string    `json:"gameCode"`
		GameStatusText    string    `json:"gameStatusText"`
		GameStatus        int       `json:"gameStatus"`
		RegulationPeriods int       `json:"regulationPeriods"`
		Period            int       `json:"period"`
		GameClock         string    `json:"gameClock"`
		Attendance        int       `json:"attendance"`
		Sellout           string    `json:"sellout"`
		Arena             struct {
			ArenaID       int    `json:"arenaId"`
			ArenaName     string `json:"arenaName"`
			ArenaCity     string `json:"arenaCity"`
			ArenaState    string `json:"arenaState"`
			ArenaCountry  string `json:"arenaCountry"`
			ArenaTimezone string `json:"arenaTimezone"`
		} `json:"arena"`
		Officials []struct {
			PersonID   int    `json:"personId"`
			Name       string `json:"name"`
			NameI      string `json:"nameI"`
			FirstName  string `json:"firstName"`
			FamilyName string `json:"familyName"`
			JerseyNum  string `json:"jerseyNum"`
			Assignment string `json:"assignment"`
		} `json:"officials"`
		HomeTeam BoxScoreTeam `json:"homeTeam"`
		AwayTeam BoxScoreTeam `json:"awayTeam"`
	} `json:"game"`
}

func getDayQuery(date string) (result DayQuery, err error) {
	endpoint := fmt.Sprintf("https://stats.nba.com/stats/internationalbroadcasterschedule?LeagueID=00&Season=%s&RegionID=1&Date=%s&EST=Y", "2023", date)
	resp, err := http.Get(endpoint)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
	}
	return result, err
}

func getBoxScoreQuery(gameId string) (result BoxScoreQuery, err error) {
	// Note that box scores are 400 error if the game hasn't started yet.
	endpoint := fmt.Sprintf("https://cdn.nba.com/static/json/liveData/boxscore/boxscore_%s.json", gameId)
	resp, err := http.Get(endpoint)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
	}
	return result, err
}

func getPlayByPlayQuery(gameId string) (result PlayByPlayQuery, err error) {
	endpoint := fmt.Sprintf("https://cdn.nba.com/static/json/liveData/playbyplay/playbyplay_%s.json", gameId)
	resp, err := http.Get(endpoint)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
	}
	err = json.Unmarshal(body, &result)
	if err != nil { // Parse []byte to the go struct pointer
	}
	return result, err
}

func loadDayQuery(date string, c chan DayQuery) {
	day, _ := getDayQuery(date)
	c <- day
}

func loadPlayByPlayQuery(gameId string, c chan *PlayByPlayQuery) {
	plays, err := getPlayByPlayQuery(gameId)
	if err != nil {
		c <- nil
		return
	}
	c <- &plays
}

func loadBoxScoreData(gameId string, c chan *BoxScoreQuery) {
	boxScore, err := getBoxScoreQuery(gameId)
	if err != nil {
		c <- nil
		return
	}
	c <- &boxScore
}

type GameContext struct {
	gameId          string
	homeTricode     string
	awayTricode     string
	time            string
	textView        *tview.TextView
	boxScore        *BoxScoreQuery
	plays           *PlayByPlayQuery
	boxScoreChannel chan *BoxScoreQuery
	playsChannel    chan *PlayByPlayQuery
	// text box?
}

type NBAContext struct {
	day        DayQuery
	dayChannel chan DayQuery
	games      []*GameContext
}

const ( // iota is reset to 0
	upcoming = iota // c0 == 0
	ongoing  = iota // c1 == 1
	finished = iota // c2 == 2
)

type GameCardRenderData struct {
	status     int
	statusText string
	header     string
	scoreText  string
	time       string
	copyText0  string
	copyText1  string
}

func (g GameContext) RenderCard() GameCardRenderData {
	var result GameCardRenderData
	result.header = fmt.Sprintf("%s@%s", g.awayTricode, g.homeTricode)
	result.time = fmt.Sprintf("%s EST", g.time)
	if g.boxScore != nil {
		result.scoreText = fmt.Sprintf("%d-%d", g.boxScore.Game.AwayTeam.Score, g.boxScore.Game.HomeTeam.Score)
		result.status = ongoing
		result.statusText = g.boxScore.Game.GameStatusText
		if strings.ToLower(result.statusText) == "final" {
			result.status = finished
			result.statusText = "Final"
		}
	} else {
		result.status = upcoming
		result.statusText = "Upcoming"
	}

	if g.plays != nil {
		homeLeadingScorerText := ""
		maxPoints := 0
		for _, player := range g.boxScore.Game.HomeTeam.Players {
			if player.Statistics.Points > maxPoints {
				maxPoints = player.Statistics.Points
				homeLeadingScorerText = fmt.Sprintf("%s. %s %d PTS %d REB %d AST", player.FirstName[0:1], player.FamilyName, player.Statistics.Points, player.Statistics.ReboundsTotal, player.Statistics.Assists)
			}
		}
		awayLeadingScorerText := ""
		maxPoints = 0
		for _, player := range g.boxScore.Game.AwayTeam.Players {
			if player.Statistics.Points > maxPoints {
				maxPoints = player.Statistics.Points
				awayLeadingScorerText = fmt.Sprintf("%s. %s %d PTS %d REB %d AST", player.FirstName[0:1], player.FamilyName, player.Statistics.Points, player.Statistics.ReboundsTotal, player.Statistics.Assists)
			}
		}
		result.copyText0 = awayLeadingScorerText
		result.copyText1 = homeLeadingScorerText
	}
	return result
}

// TODO render focused view data with box score, detailed stats etc

func main() {
	currentDate := time.Now().Format("01/02/2006")
	fmt.Printf("Current date: %s\n", currentDate)
	var nbaContext NBAContext
	nbaContext.dayChannel = make(chan DayQuery)
	go loadDayQuery(currentDate, nbaContext.dayChannel)
	nbaContext.day = <-nbaContext.dayChannel
	nbaContext.games = []*GameContext{}
	for _, gameMd := range nbaContext.day.ResultSets[1].CompleteGameList {
		gameId := gameMd.GameID
		var gameContext GameContext
		gameContext.gameId = gameId
		gameContext.homeTricode = gameMd.HtAbbreviation
		gameContext.awayTricode = gameMd.VtAbbreviation
		gameContext.time = gameMd.Time
		gameContext.boxScoreChannel = make(chan *BoxScoreQuery)
		gameContext.playsChannel = make(chan *PlayByPlayQuery)
		go loadBoxScoreData(gameId, gameContext.boxScoreChannel)
		go loadPlayByPlayQuery(gameId, gameContext.playsChannel)
		nbaContext.games = append(nbaContext.games, &gameContext)
	}
	for _, gameContext := range nbaContext.games {
		gameContext.boxScore = <-gameContext.boxScoreChannel
		gameContext.plays = <-gameContext.playsChannel
	}

	renderCellText := func(card GameCardRenderData) string {
		statusColor := "red"
		if card.status == ongoing {
			statusColor = "green"
		}
		gameTime := ""
		if card.status == upcoming {
			gameTime = fmt.Sprintf("%s\n", card.time)
		}
		return fmt.Sprintf("[white::b]%s\n%s[%s::b]%s[white::-]\n%s\n%s", card.scoreText, gameTime, statusColor, card.statusText, card.copyText0, card.copyText1)
	}

	renderCell := func(card GameCardRenderData, child *tview.TextView) *tview.Frame {
		frame := tview.NewFrame(child).
			SetBorders(1, 1, 1, 1, 1, 1)
		frame.SetBorder(true).SetTitle(card.header)
		return frame
	}

	app := tview.NewApplication()
	grid := tview.NewGrid()

	for index, gameContext := range nbaContext.games {
		r := index / 5
		c := index % 5
		card := gameContext.RenderCard()
		textView := tview.NewTextView().
			SetDynamicColors(true).
			SetTextAlign(tview.AlignCenter).
			SetText(renderCellText(card))
		textView.SetChangedFunc(func() {
			app.Draw()
		})
		gameContext.textView = textView
		grid.AddItem(renderCell(card, textView), r, c, 1, 1, 0, 0, false)
	}

	updateElements := func() {
		for _, gameContext := range nbaContext.games {
			gameId := gameContext.gameId
			go loadBoxScoreData(gameId, gameContext.boxScoreChannel)
			go loadPlayByPlayQuery(gameId, gameContext.playsChannel)
		}
		for _, gameContext := range nbaContext.games {
			gameContext.boxScore = <-gameContext.boxScoreChannel
			gameContext.plays = <-gameContext.playsChannel
		}
		for _, gameContext := range nbaContext.games {
			gameContext.textView.SetText(renderCellText(gameContext.RenderCard()))
		}
	}

	// TODO only do this for visible elements
	go (func() {
		// this could read state, maybe
		for true {
			time.Sleep(5 * time.Second)
			updateElements()
		}
	})()

	headerText := fmt.Sprintf("Welcome to NBA TUI!\n%s", currentDate)
	header := tview.NewFrame(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(headerText)).SetBorders(1, 1, 1, 1, 1, 1)
	header.SetBorder(true).SetTitle("NBA TUI")

	homeView := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(header, 0, 1, true).AddItem(grid, 0, 5, true)
	renderDetailedView := func(gameContext *GameContext) *tview.Flex {
		renderPlayerBoxScore := func(teamData *BoxScoreTeam) *tview.Table {
			table := tview.NewTable().SetBorders(true)
			table.SetCell(0, 0, tview.NewTableCell(fmt.Sprintf("%s                     ", teamData.TeamTricode)))
			for i, colName := range []string{"MIN", "PTS", "TREB", "AST", "   FG", "  3PT", "   FT", "OREB", "DREB", "STL", "BLK", "TO", "PF", "+/-"} {
				table.SetCell(0, i+1, tview.NewTableCell(colName).SetAlign(tview.AlignRight))
			}
			for r, player := range teamData.Players {
				table.SetCell(r+1, 0, tview.NewTableCell(fmt.Sprintf("%s. %s #%s %s", player.FirstName[0:1], player.FamilyName, player.JerseyNum, player.Position)))
				minutes, err := strconv.Atoi(player.Statistics.MinutesCalculated[2:4])
				if err == nil {
					table.SetCell(r+1, 1, tview.NewTableCell(fmt.Sprintf("%d", minutes)))
				}
				table.SetCell(r+1, 2, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.Points)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 3, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.ReboundsTotal)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 4, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.Assists)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 5, tview.NewTableCell(fmt.Sprintf("%d-%d", player.Statistics.FieldGoalsMade, player.Statistics.FieldGoalsAttempted)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 6, tview.NewTableCell(fmt.Sprintf("%d-%d", player.Statistics.ThreePointersMade, player.Statistics.ThreePointersAttempted)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 7, tview.NewTableCell(fmt.Sprintf("%d-%d", player.Statistics.FreeThrowsMade, player.Statistics.FreeThrowsAttempted)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 8, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.ReboundsOffensive)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 9, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.ReboundsDefensive)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 10, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.Steals)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 11, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.Blocks)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 12, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.Turnovers)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 13, tview.NewTableCell(fmt.Sprintf("%d", player.Statistics.FoulsPersonal)).SetAlign(tview.AlignRight))
				table.SetCell(r+1, 14, tview.NewTableCell(fmt.Sprintf("%0.f", player.Statistics.PlusMinusPoints)).SetAlign(tview.AlignRight))
			}
			return table
		}
		renderTeamBoxScores := func() *tview.Table {
			table := tview.NewTable().SetBorders(true)
			renderTeamRow := func(row int, teamData *BoxScoreTeam) {
				table.SetCell(row, 1, tview.NewTableCell(fmt.Sprintf("%d", teamData.Statistics.Points)).SetAlign(tview.AlignRight))
				table.SetCell(row, 2, tview.NewTableCell(fmt.Sprintf("%d-%d", teamData.Statistics.FieldGoalsMade, teamData.Statistics.FieldGoalsAttempted)).SetAlign(tview.AlignRight))
				table.SetCell(row, 3, tview.NewTableCell(fmt.Sprintf("%d-%d", teamData.Statistics.ThreePointersMade, teamData.Statistics.ThreePointersAttempted)).SetAlign(tview.AlignRight))
				table.SetCell(row, 4, tview.NewTableCell(fmt.Sprintf("%d", teamData.Statistics.ReboundsOffensive)).SetAlign(tview.AlignRight))
				table.SetCell(row, 5, tview.NewTableCell(fmt.Sprintf("%d", teamData.Statistics.ReboundsDefensive)).SetAlign(tview.AlignRight))
				table.SetCell(row, 6, tview.NewTableCell(fmt.Sprintf("%d", teamData.Statistics.ReboundsTotal)).SetAlign(tview.AlignRight))
				table.SetCell(row, 7, tview.NewTableCell(fmt.Sprintf("%d", teamData.Statistics.Assists)).SetAlign(tview.AlignRight))
				table.SetCell(row, 8, tview.NewTableCell(fmt.Sprintf("%d", teamData.Statistics.Turnovers)).SetAlign(tview.AlignRight))
			}
			table.SetCell(0, 0, tview.NewTableCell(""))
			table.SetCell(1, 0, tview.NewTableCell(gameContext.awayTricode))
			table.SetCell(2, 0, tview.NewTableCell(gameContext.homeTricode))
			for i, colName := range []string{"PTS", "FG", "3PT", "ORB", "DRB", "TRB", "AST", "TO"} {
				table.SetCell(0, i+1, tview.NewTableCell(colName).SetAlign(tview.AlignRight))
			}
			renderTeamRow(1, &gameContext.boxScore.Game.AwayTeam)
			renderTeamRow(2, &gameContext.boxScore.Game.HomeTeam)
			return table
		}
		renderPlayByPlay := func() *tview.TextView { // TODO render incrementally, don't pull entire
			p := tview.NewTextView()
			s := ""
			for _, action := range gameContext.plays.Game.Actions {
				s += action.Description + "\n"
			}
			p.SetText(s)
			p.ScrollToEnd()
			return p
		}
		teamStatsTable := renderTeamBoxScores()
		teamStatsFlex := tview.NewFlex().AddItem(tview.NewBox(), 0, 1, true).AddItem(teamStatsTable, 0, 1, true).AddItem(tview.NewBox(), 0, 1, true)
		top := tview.NewFrame(teamStatsFlex).
			AddText(fmt.Sprintf("%s %s at %s %s",
				gameContext.boxScore.Game.AwayTeam.TeamCity,
				gameContext.boxScore.Game.AwayTeam.TeamName,
				gameContext.boxScore.Game.HomeTeam.TeamCity,
				gameContext.boxScore.Game.HomeTeam.TeamName), true, tview.AlignCenter, tcell.ColorWhite).
			AddText(fmt.Sprintf("%s, %s", gameContext.boxScore.Game.Arena.ArenaName, gameContext.boxScore.Game.Arena.ArenaCity), true, tview.AlignCenter, tcell.ColorWhite).
			AddText(fmt.Sprintf("%s", gameContext.boxScore.Game.GameTimeUTC), true, tview.AlignCenter, tcell.ColorWhite).
			AddText(gameContext.boxScore.Game.GameStatusText, true, tview.AlignCenter, tcell.ColorWhite)
		awayPlayerTable := renderPlayerBoxScore(&gameContext.boxScore.Game.AwayTeam)
		homePlayerTable := renderPlayerBoxScore(&gameContext.boxScore.Game.HomeTeam)
		playByPlay := renderPlayByPlay()
		/*
			const (
				home = iota
				away = iota
				pbp  = iota
			)
			whichView := home
		*/
		frame := tview.NewFrame(awayPlayerTable)
		app.SetFocus(awayPlayerTable)
		menuText := tview.NewTextView().SetText("[1] Away [2] Home [3] Play-By-Play")
		menu := tview.NewFrame(menuText) // TODO highlight selected
		view := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(top, 0, 4, true).
			AddItem(menu, 0, 1, true).
			AddItem(frame, 0, 10, true)
		view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == '1' {
				frame.SetPrimitive(awayPlayerTable) // TODO: rerender then async fetch
				app.SetFocus(awayPlayerTable)
			} else if event.Rune() == '2' {
				frame.SetPrimitive(homePlayerTable) // TODO: rerender then fetch
				app.SetFocus(homePlayerTable)
			} else if event.Rune() == '3' {
				frame.SetPrimitive(playByPlay)
				app.SetFocus(playByPlay)
			}
			//fmt.Println("hello ")
			return event
		})
		return view
	}

	focusIndex := 0
	const (
		homepage = iota
		detailed = iota
	)
	currentView := homepage

	rotateFocus := func(i int) {
		mod := func(a, b int) int {
			return (a%b + b) % b
		}
		focusIndex += i
		focusIndex = mod(focusIndex, len(nbaContext.games))
		app.SetFocus(nbaContext.games[focusIndex].textView)
	}

	app.SetRoot(homeView, true)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Name() == "Tab" {
			if currentView == homepage {
				rotateFocus(1)
			}
		} else if event.Name() == "Enter" {
			if currentView == homepage {
				if nbaContext.games[focusIndex].boxScore != nil {
					view := renderDetailedView(nbaContext.games[focusIndex])
					app.SetRoot(view, true)
					app.SetFocus(view)
					currentView = detailed
				}
			}
		} else if event.Name() == "Esc" || event.Rune() == 'q' {
			if currentView == detailed {
				app.SetRoot(homeView, true)
				app.SetFocus(nbaContext.games[focusIndex].textView)
				currentView = homepage
			} /*else if currentView == homepage {
				// TODO: consider not using ESC to exit program
				//return tcell.NewEventKey(tcell.KeyCtrlC, ' ', tcell.ModNone)
				app.Stop()
				return nil
			}*/
		}

		if currentView == homepage {
			if event.Name() == "Left" || event.Rune() == 'h' {
				rotateFocus(-1)
				return nil
			} else if event.Name() == "Right" || event.Rune() == 'l' {
				rotateFocus(1)
				return nil
			}
			if event.Name() == "Up" || event.Rune() == 'k' {
				rotateFocus(-5)
				return nil
			} else if event.Name() == "Down" || event.Rune() == 'j' {
				rotateFocus(5)
				return nil
			}
		}
		return event
	})

	if len(nbaContext.games) > 0 {
		app.SetFocus(nbaContext.games[focusIndex].textView)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
	// TODO: track total # of requests and try to minimize (ie don't keep requesting after game ends, etc)
}
