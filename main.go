package main

import "github.com/projects/threaded-company-simulation/factory"

func main() {
	sync := make(chan factory.Job, 10)
	factory.Ceo(sync)
}
