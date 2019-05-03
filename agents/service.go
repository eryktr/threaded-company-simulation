package agents

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/projects/threaded-company-simulation/config"
)

const ADDITION = 0
const MULTIPLICATION = 1

type ReportChanelWriteOp struct {
	report BreakdownReport
	result chan bool
}

type ReportChannelReadOp struct {
	report chan BreakdownReport
	result chan bool
}

type ServiceWorker struct {
	Id            int
	ReportChannel chan ReportChannelReadOp
	MulltMachines []*MultiplicationMachine
	AddMachines   []*AdditionMachine
	FixChannel    chan FixReport
	Logger        chan string
}

type BreakdownReport struct {
	MachineType     int
	MachineNumber   int
	BreakdownNumber int
}

type FixReport struct {
	MachineType     int
	MachineNumber   int
	BreakdownNumber int
}

type Service struct {
	ReportWrite chan ReportChanelWriteOp
	ReportRead  chan ReportChannelReadOp
	FixChannel  chan FixReport
	Reports     []BreakdownReport
	ReportCache []BreakdownReport
	Logger      chan string
}

func (s *Service) Run() {
	for {
		select {
		case report := <-s.ReportWrite:
			newReport := report.report
			alreadyReported := false
			for i := 0; i < len(s.ReportCache); i++ {
				report := s.ReportCache[i]
				mtype := newReport.MachineType
				index := newReport.MachineNumber
				breakdown := newReport.BreakdownNumber
				if report.MachineType == mtype && report.MachineNumber == index && report.BreakdownNumber == breakdown {
					alreadyReported = true
				}
			}
			if !alreadyReported {
				s.Reports = append(s.Reports, report.report)
				s.ReportCache = append(s.ReportCache, newReport)
				s.Logger <- fmt.Sprintf("SERVICE: RECEIVED NEW BREAKDOWN INFORMATION ABOUT MACHINE OF TYPE %d and ID %d\n",
					newReport.MachineType, newReport.MachineNumber)
			}
			report.result <- true

		case request := <-maybeReadReport(len(s.Reports) > 0, s.ReportRead):
			response := s.Reports[0]
			s.Reports = s.Reports[1:]
			request.report <- response
			// This line caused a deadlock - no idea why
			//request.result <- true

		case fix := <-s.FixChannel:
			mtype := fix.MachineType
			index := fix.MachineNumber
			breakdown := fix.MachineNumber
			for i := 0; i < len(s.ReportCache); i++ {
				report := s.ReportCache[i]
				if report.MachineType == mtype && report.MachineNumber == index && report.BreakdownNumber == breakdown {
					s.ReportCache = append(s.ReportCache[:i], s.ReportCache[i+1:]...)
				}
			}
		}
	}
}

func (sw *ServiceWorker) Sleep() {
	randnum := rand.Float32()
	time.Sleep(time.Duration(randnum * config.AVERAGE_SERVICE_WORKER_DELAY))
}

func (sw *ServiceWorker) Run() {
	for {
		responseChannel := make(chan BreakdownReport)
		resultChannel := make(chan bool)
		request := ReportChannelReadOp{
			report: responseChannel,
			result: resultChannel,
		}
		sw.ReportChannel <- request
		task := <-request.report
		mtype := task.MachineType
		id := task.MachineNumber
		if mtype == 0 {
			sw.Logger <- fmt.Sprintf("SERVICE WORKER %d: FIXING ADDITION MACHINE %d \n",
				sw.Id, id)
			sw.AddMachines[id].FixMe <- true
			sw.Logger <- fmt.Sprintf("SERVICE WORKER %d: FIXED ADDITION MACHINE %d \n",
				sw.Id, id)
		} else {
			sw.Logger <- fmt.Sprintf("SERVICE WORKER %d: FIXING MULTIPLICATION MACHINE %d \n",
				sw.Id, id)
			sw.MulltMachines[id].FixMe <- true
			sw.Logger <- fmt.Sprintf("SERVICE WORKER %d: FIXED MULTIPLICATION MACHINE %d \n",
				sw.Id, id)
		}
		fix := FixReport{
			BreakdownNumber: task.BreakdownNumber,
			MachineNumber:   task.MachineNumber,
			MachineType:     task.MachineType,
		}
		sw.FixChannel <- fix
	}
}

func maybeReadReport(expression bool, c chan ReportChannelReadOp) chan ReportChannelReadOp {
	if expression {
		return c
	} else {
		return nil
	}
}
