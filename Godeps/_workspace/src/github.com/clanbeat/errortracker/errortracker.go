package errortracker

import (
	"errors"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/getsentry/raven-go"
	"log"
)

type ErrorTracker interface {
	Error(err error)
	Message(msg string)
	Recover(interface{})
	Wait()
}

type Tracker struct {
	Client *raven.Client
	Env    string
}

func New(sentryDSN, env string) (*Tracker, error) {
	if len(sentryDSN) == 0 {
		return &Tracker{Env: env}, nil
	}
	client, err := raven.NewClient(sentryDSN, nil)
	if err != nil {
		return nil, err
	}
	return &Tracker{Client: client, Env: env}, nil
}

func (t *Tracker) Wait() {
	if t.Client != nil {
		t.Client.Wait()
	}
}

func (t *Tracker) Error(err error) {
	logError(err)
	if t.Client != nil {
		t.Client.CaptureError(err, nil)
	}
}

func (t *Tracker) ErrorAndWait(err error) {
	logError(err)
	if t.Client != nil {
		t.Client.CaptureErrorAndWait(err, nil)
	}
}

func (t *Tracker) Message(msg string) {
	if t.Client != nil {
		t.Client.CaptureMessage(msg, nil)
	}
}

func (t *Tracker) Recover(msg interface{}) {
	var err error
	switch t := msg.(type) {
	case string:
		err = errors.New(t)
	case error:
		err = t
	default:
		err = errors.New("Unknown error")
	}
	t.ErrorAndWait(err)
}

func logError(err error) {
	log.Println("[error_tracker]:", err)
}
