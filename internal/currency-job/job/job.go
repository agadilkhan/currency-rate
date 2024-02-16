package job

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-job/service"
	"log"
	"time"
)

type Job struct {
	service  service.UseCase
	interval time.Duration
}

func New(srvc service.UseCase, interval time.Duration) *Job {
	return &Job{
		service:  srvc,
		interval: interval,
	}
}

func (j *Job) Run(ctx context.Context) {

	ticker := time.NewTimer(j.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			startTime := time.Now()

			err := j.service.Update(ctx)
			if err != nil {
				log.Printf("failed to Update err: %v", err)
				return
			}

			elapsedTime := time.Since(startTime)

			ticker.Reset(j.interval - elapsedTime)
		}
	}
}
