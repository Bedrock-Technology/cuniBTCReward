package crons

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"
)

type ScanCron struct {
	Cron *cron.Cron
}

func New() *ScanCron {
	scanCron := &ScanCron{}
	scanCron.Cron = cron.New(cron.WithLocation(time.UTC), cron.WithChain(cron.SkipIfStillRunning(cron.DiscardLogger)), cron.WithSeconds())
	return scanCron
}

func (sc *ScanCron) AddSpec(spec string, job cron.Job) (cron.EntryID, error) {
	wrappedJob := cron.NewChain(cron.SkipIfStillRunning(cron.DiscardLogger)).Then(job)
	return sc.Cron.AddJob(spec, wrappedJob)
}

func (sc *ScanCron) Start() {
	sc.Cron.Run()
}

func (sc *ScanCron) Run() {
	sc.Cron.Start()
}

func (sc *ScanCron) StopContext() context.Context {
	return sc.Cron.Stop()
}

func (sc *ScanCron) Stop() {
	ctx := sc.Cron.Stop()
	<-ctx.Done()
}

func (sc *ScanCron) Remove(id cron.EntryID) {
	sc.Cron.Remove(id)
}

func (sc *ScanCron) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return sc.Cron.AddFunc(spec, cmd)
}
