package factory

import (
	"time"

	"github.com/projects/threaded-company-simulation/config"
)

func sleep_failure() {
	time.Sleep(time.Duration(config.FAILURE_DELAY_TIME))
}
