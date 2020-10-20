package hawk

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
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
func readSourceCode(reader io.Reader, targetLine int) ([3]SourceCode, error) {
	var res [3]SourceCode
	lines := make([]string, 3)
	scanner := bufio.NewScanner(reader)
	i := 1
	for scanner.Scan() {
		switch i {
		case targetLine - 1:
			lines[0] = scanner.Text()
		case targetLine:
			lines[1] = scanner.Text()
		case targetLine + 1:
			lines[2] = scanner.Text()
		}
		if i == targetLine+1 {
			break
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		return res, err
	}

	res = [3]SourceCode{
		{
			LineNumber: targetLine - 1,
			Content:    strings.Trim(lines[0], "\t"),
		},
		{
			LineNumber: targetLine,
			Content:    strings.Trim(lines[1], "\t"),
		},
		{
			LineNumber: targetLine + 1,
			Content:    strings.Trim(lines[2], "\t"),
		},
	}

	return res, nil
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
			Timestamp: strconv.Itoa(int(time.Now().Unix())),
		},
	}

	report.Payload.Backtrace = getBacktrace(1)
	if len(report.Payload.Backtrace) == 0 {
		return ErrEmptyBacktrace
	}

	for i, bt := range report.Payload.Backtrace {
		file, err := os.Open(bt.File)
		if err != nil {
			log.Printf("failed to open file %s: %s", bt.File, err.Error())
			continue
		}

		report.Payload.Backtrace[i].SourceCode, err = readSourceCode(file, bt.Line)
		if err != nil {
			log.Printf("failed to read file %s: %s", bt.File, err.Error())
			continue
		}
		file.Close()
	}

	return c.proceedReport(&report)
}
