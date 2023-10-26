package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	//"net"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"net/http"
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

func getDayQuery() (result DayQuery, err error) {
	endpoint := fmt.Sprintf("https://stats.nba.com/stats/internationalbroadcasterschedule?LeagueID=00&Season=%s&RegionID=1&Date=%s&EST=Y", "2023", "10/25/2023")
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Cannot unmarshal JSON")
	}
	return result, err
}

func getBoxScoreQuery(gameId string) (result BoxScoreQuery, err error) {
	endpoint := fmt.Sprintf("https://cdn.nba.com/static/json/liveData/boxscore/boxscore_%s.json", gameId)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Cannot unmarshal JSON")
	}
	return result, err

}

func getPlayByPlayQuery(gameId string) (result PlayByPlayQuery, err error) {
	endpoint := fmt.Sprintf("https://cdn.nba.com/static/json/liveData/playbyplay/playbyplay_%s.json", gameId)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Cannot unmarshal JSON")
	}
	return result, err
}

func main() {
	flexItems := [3]*tview.Flex{
		tview.NewFlex().SetDirection(tview.FlexRow),
		tview.NewFlex().SetDirection(tview.FlexRow),
		tview.NewFlex().SetDirection(tview.FlexRow),
	}
	app := tview.NewApplication()
	day, err := getDayQuery()
	if err != nil {
		log.Fatal("Failed to query day")
	}
	for index, md := range day.ResultSets[1].CompleteGameList {
		gameId := md.GameID
		boxScore, err := getBoxScoreQuery(gameId)
		if err != nil {
			fmt.Println("failed to query box score")
			continue
		}
		header := fmt.Sprintf("%s@%s", boxScore.Game.AwayTeam.TeamTricode, boxScore.Game.HomeTeam.TeamTricode)
		headToHeadScore := fmt.Sprintf("%d-%d", boxScore.Game.AwayTeam.Score, boxScore.Game.HomeTeam.Score)
		//fmt.Printf("Status: %s %d\n", boxScore.Game.GameStatusText, boxScore.Game.GameStatus)
		frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
			SetBorders(2, 2, 2, 2, 4, 4).
			AddText(headToHeadScore, true, tview.AlignCenter, tcell.ColorWhite).
			AddText(boxScore.Game.GameStatusText, true, tview.AlignCenter, tcell.ColorRed).
			AddText("Footer middle", false, tview.AlignCenter, tcell.ColorGreen).
			AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen)
		frame.Box.SetBorder(true).SetTitle(header)
		currentRow := index / 5
		flexItems[currentRow].AddItem(frame, 0, 1, false)

		/*
			pbp, err := getPlayByPlayQuery(gameId)
			if err != nil {
				log.Fatal("Failed to query day")
			}
				for _, action := range pbp.Game.Actions {
					fmt.Println(action.Description)
				}
		*/
	}
	flex := tview.NewFlex()
	flex.AddItem(flexItems[0], 0, 1, false)
	flex.AddItem(flexItems[1], 0, 1, false)
	flex.AddItem(flexItems[2], 0, 1, false)
	/*
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	*/
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
