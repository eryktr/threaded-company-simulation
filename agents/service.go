package agents

import "fmt"

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
	reportChannel chan BreakdownReport
	fixChannel    chan FixReport
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
				mtype := report.MachineType
				index := report.MachineNumber
				breakdown := report.BreakdownNumber
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

		case request := <-maybeReadReport(len(s.ReportCache) > 0, s.ReportRead):
			response := s.Reports[0]
			s.Reports = s.Reports[1:]
			request.report <- response
			request.result <- true
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

func maybeReadReport(expression bool, c chan ReportChannelReadOp) chan ReportChannelReadOp {
	if expression {
		return c
	} else {
		return nil
	}
}
