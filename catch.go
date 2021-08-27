package hawk

import (
	"bufio"
	json "encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/mailru/easyjson"
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
		res = append(res, Backtrace{
			File:     frame.File,
			Line:     frame.Line,
			Function: frame.Function,
		})

	}
	return res
}

// readSourceCode reads the lines of code that caused the error.
func (c *Catcher) readSourceCode(reader io.Reader, targetLine int) ([]SourceCode, error) {
	var res []SourceCode
	lines := []string{}
	scanner := bufio.NewScanner(reader)
	idx := 1
	delta := c.options.SourceCodeLines
	for scanner.Scan() {
		if idx == (targetLine - delta) {
			lines = append(lines, scanner.Text())
			delta--
		}
		if idx == targetLine+c.options.SourceCodeLines {
			break
		}
		idx++
	}
	if err := scanner.Err(); err != nil {
		return res, err
	}

	res = []SourceCode{}
	delta = c.options.SourceCodeLines
	for i := range lines {
		res = append(res, SourceCode{
			LineNumber: targetLine - delta,
			Content:    lines[i],
		})
		delta--
	}

	return res, nil
}

// Catch takes error object and additional options to pass to the catchWithPayload
func (c *Catcher) Catch(err error, opts ...HawkAdditionalParams) error {
	if err == nil {
		return nil
	}

	h := &Payload{
		Title:          err.Error(),
		Type:           ManualType,
		CatcherVersion: VERSION,
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated *Payload as the argument
		opt(h)
	}

	// return the modified Payload instance
	return c.catchWithPayload(*h)
}

// catchWithPayload creates ErrorReport for provided error, collects backtrace and sends data to Hawk.
func (c *Catcher) catchWithPayload(payload Payload) error {
	if payload.User.isEmpty() {
		payload.User = c.options.AffectedUser
	}

	// add integrationID to context if Debug is enabled
	if c.integrationID != "" && c.options.Debug {
		var context map[string]interface{}
		if payload.Context == nil {
			payload.Context = easyjson.RawMessage(`{}`)
		}
		err := json.Unmarshal(payload.Context, &context)
		if err == nil {
			context["integrationID"] = c.integrationID
			newContext, err := json.Marshal(context)
			if err == nil {
				payload.Context = newContext
			}
		}
	}

	report := ErrorReport{
		Token:       c.options.AccessToken,
		CatcherType: CatcherType,
		Payload:     payload,
	}

	report.Payload.Backtrace = getBacktrace(2)
	if len(report.Payload.Backtrace) == 0 {
		return ErrEmptyBacktrace
	}

	if c.options.SourceCodeEnabled {
		for i, bt := range report.Payload.Backtrace {
			file, err := os.Open(bt.File)
			if err != nil {
				log.Printf("failed to open file %s: %s", bt.File, err.Error())
				break
			}

			report.Payload.Backtrace[i].SourceCode, err = c.readSourceCode(file, bt.Line)
			if err != nil {
				log.Printf("failed to read file %s: %s", bt.File, err.Error())
				file.Close()
				continue
			}
			file.Close()
		}
	}

	c.errorsCh <- report

	return nil
}

func WithContext(context interface{}) HawkAdditionalParams {
	return func(h *Payload) {
		if context != nil {
			b, err := json.Marshal(&context)
			if err == nil {
				h.Context = b
			}
		}
	}
}

func WithUser(user AffectedUser) HawkAdditionalParams {
	return func(h *Payload) {
		h.User = user
	}
}

func WithRelease(release string) HawkAdditionalParams {
	return func(h *Payload) {
		h.Release = release
	}
}
