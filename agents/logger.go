package agents

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"
)

func RunLogger() {
	for {
		select {
		case read := <-LogChannel:
			if config.MODE == config.VERBOSE {
				fmt.Println(read)
			}
		}
	}
}
