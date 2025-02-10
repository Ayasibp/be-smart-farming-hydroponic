package middleware

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/go-co-op/gocron"
)

type CronJob interface {
	CreateAggregationEachMonth()
}

type cronJob struct {
	aggregateService service.AggregationService
}

type CronJobConfig struct {
	AggregateService service.AggregationService
}

func NewCorn(config CronJobConfig) CronJob {
	return &cronJob{
		aggregateService: config.AggregateService,
	}
}

func (c cronJob) CreateAggregationEachMonth() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Cron("0 0 1 * *").Do(func() {

	})

	scheduler.StartAsync()
}
