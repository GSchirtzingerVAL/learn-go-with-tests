package poker_test

import (
	"fmt"
	"testing"
	"time"

	poker "gerrod.com/http-server"
)

var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyBlindAlerter = &poker.SpyBlindAlerter{}

func TestGame_Start(t *testing.T) {

	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []poker.ScheduledAlert{
			poker.NewScheduledAlert(0*time.Second, 100),
			poker.NewScheduledAlert(10*time.Minute, 200),
			poker.NewScheduledAlert(20*time.Minute, 300),
			poker.NewScheduledAlert(30*time.Minute, 400),
			poker.NewScheduledAlert(40*time.Minute, 500),
			poker.NewScheduledAlert(50*time.Minute, 600),
			poker.NewScheduledAlert(60*time.Minute, 800),
			poker.NewScheduledAlert(70*time.Minute, 1000),
			poker.NewScheduledAlert(80*time.Minute, 2000),
			poker.NewScheduledAlert(90*time.Minute, 4000),
			poker.NewScheduledAlert(100*time.Minute, 8000),
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			poker.NewScheduledAlert(0*time.Second, 100),
			poker.NewScheduledAlert(12*time.Minute, 200),
			poker.NewScheduledAlert(24*time.Minute, 300),
			poker.NewScheduledAlert(36*time.Minute, 400),
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			poker.AssertAlerts(t, blindAlerter, i, want)

		})
	}
}

func TestGame_Finish(t *testing.T) {

	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}
