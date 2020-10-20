package hawk

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// ErrEmptyBacktrace is returned if getBacktrace collected empty backtrace.
var ErrEmptyBacktrace = errors.New("failed to collect backtrace")

// getBacktrace collects backtrace data.
func getBacktrace(toSkip int) []Backtrace {
	pc := make([]uintptr, 64)
	var numFrames int
	for {
		numFrames = runtime.Callers(toSkip+2, pc)
		if numFrames < len(pc) {
			break
		}
	}

	res := []Backtrace{}
	frames := runtime.CallersFrames(pc[:numFrames])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if strings.Contains(frame.File, "runtime") {
			break
		}

		res = append(res, Backtrace{
			File: frame.File,
			Line: frame.Line,
		})

	}
	return res
}

// readSourceCode reads the lines of code that caused the error.
func (b *Backtrace) readSourceCode(reader io.Reader) error {
	lines := make([]string, 3)
	scanner := bufio.NewScanner(reader)
	i := 1
	for scanner.Scan() {
		switch i {
		case b.Line - 1:
			lines[0] = scanner.Text()
		case b.Line:
			lines[1] = scanner.Text()
		case b.Line + 1:
			lines[2] = scanner.Text()
		}
		if i == b.Line+1 {
			break
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	b.SourceCode = [3]SourceCode{
		{
			LineNumber: b.Line - 1,
			Content:    lines[0],
		},
		{
			LineNumber: b.Line,
			Content:    lines[1],
		},
		{
			LineNumber: b.Line + 1,
			Content:    lines[2],
		},
	}

	return nil
}

// Catch creates ErrorReport for provided error, collects backtrace and sends
// data to Hawk.
func (c *Catcher) Catch(err error) error {
	if err == nil {
		return nil
	}

	report := ErrorReport{
		Token:       c.accessToken,
		CatcherType: CatcherType,
		Payload: Payload{
			Title:     err.Error(),
			Timestamp: time.Now().String(),
		},
	}

	backtraceList := getBacktrace(1)
	if len(backtraceList) == 0 {
		return ErrEmptyBacktrace
	}

	file, err := os.Open(backtraceList[0].File)
	if err != nil {
		log.Printf("failed to open file %s: %s", backtraceList[0].File, err.Error())
	}

	for i, bt := range backtraceList {
		if (i != 0) && (bt.File != backtraceList[i-1].File) {
			file.Close()
			file, err = os.Open(bt.File)
			if err != nil {
				log.Printf("failed to open file %s: %s", bt.File, err.Error())
				continue
			}
		}

		err = bt.readSourceCode(file)
		if err != nil {
			log.Printf("failed to read file %s: %s", bt.File, err.Error())
			continue
		}
	}
	file.Close()

	report.Payload.Backtrace = backtraceList

	return c.proceedReport(&report)
}
