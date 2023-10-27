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
	"reflect"
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
		HomeTeam struct {
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
				BiggestLeadScore             string  `json:"biggestLeadScore"`
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
		} `json:"homeTeam"`
		AwayTeam struct {
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
		} `json:"awayTeam"`
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
	currentDate := "10/26/2023" //time.Now().Format("01/02/2006")
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

	go (func() {
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
		table := tview.NewTable().SetBorders(true)
		if gameContext.boxScore != nil {
			table.SetCell(0, 0, tview.NewTableCell("Player"))
			rval := reflect.ValueOf(gameContext.boxScore.Game.AwayTeam.Players[0].Statistics)
			for i := 0; i < rval.Type().NumField(); i++ {
				fname := rval.Type().Field(i).Name
				table.SetCell(0, i+1, tview.NewTableCell(fname))
			}
			for r, player := range gameContext.boxScore.Game.AwayTeam.Players {
				table.SetCell(r+1, 0, tview.NewTableCell(fmt.Sprintf("%s. %s", player.FirstName[0:1], player.FamilyName)))
				rval := reflect.ValueOf(player.Statistics)
				for i := 0; i < rval.Type().NumField(); i++ {
					f := reflect.Indirect(rval).Field(i)
					table.SetCell(r+1, i+1, tview.NewTableCell(fmt.Sprintf("%s", f)))
				}
			}
		}
		return tview.NewFlex().SetDirection(tview.FlexRow).AddItem(tview.NewTextView().SetText(gameContext.gameId), 0, 1, true).AddItem(table, 0, 4, true)
	}

	focusIndex := 0
	currentView := 0

	app.SetRoot(homeView, true)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// TODO capture arrow keys/vim keys as well
		if event.Name() == "Tab" {
			if currentView == 0 {
				focusIndex += 1
				if focusIndex >= len(nbaContext.games) {
					focusIndex = 0
				}
				app.SetFocus(nbaContext.games[focusIndex].textView)
			}
		} else if event.Name() == "Enter" {
			if currentView == 0 {
				app.SetRoot(renderDetailedView(nbaContext.games[focusIndex]), true)
				currentView = 1
			}
		} else if event.Name() == "Esc" {
			if currentView == 1 {
				app.SetRoot(homeView, true)
				app.SetFocus(nbaContext.games[focusIndex].textView)
				currentView = 0
			}
		}
		//fmt.Println(event.Name())
		return event
	})

	if len(nbaContext.games) > 0 {
		app.SetFocus(nbaContext.games[focusIndex].textView)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
