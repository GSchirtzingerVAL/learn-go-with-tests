package poker

import (
	"fmt"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func NewStubPlayerStore(scores map[string]int, winCalls []string, league []Player) *StubPlayerStore {
	return &StubPlayerStore{
		scores:   scores,
		winCalls: winCalls,
		league:   league,
	}
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}

type ScheduledAlert struct {
	at     time.Duration
	amount int
}

func NewScheduledAlert(at time.Duration, amount int) ScheduledAlert {
	return ScheduledAlert{at, amount}
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{at, amount})
}

func AssertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertAlerts(t testing.TB, blindAlerter *SpyBlindAlerter, i int, want ScheduledAlert) {
	t.Helper()

	if len(blindAlerter.alerts) <= i {
		t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
	}

	got := blindAlerter.alerts[i]
	AssertScheduledAlert(t, got, want)

}
