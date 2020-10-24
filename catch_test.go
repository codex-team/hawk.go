package hawk

import (
	"os"
	"reflect"
	"sync"
	"testing"
)

// TestBacktrace tests collecting backtrace and getting source code.
func TestBacktrace(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		bt := getBacktrace(0)
		if len(bt) == 0 {
			t.Fatalf("received empty backtrace")
		}
		file, err := os.Open(bt[0].File)
		if err != nil {
			t.Fatalf("cannot open file %s: %s", bt[0].File, err.Error())
		}
		defer file.Close()
		catcher := &Catcher{SourceCodeLines: 3}
		bt[0].SourceCode, err = catcher.readSourceCode(file, bt[0].Line)
		if err != nil {
			t.Fatalf("failed to read source code: %s", err.Error())
		}

		pwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("failed to get pwd: %s", err.Error())
		}

		targetLine := 13
		expected := Backtrace{
			File:     pwd + "/catch_test.go",
			Line:     targetLine,
			Function: "github.com/codex-team/hawk%2ego.TestBacktrace.func1",
			SourceCode: []SourceCode{
				{
					LineNumber: targetLine - 1,
					Content:    "\tt.Run(\"simple case\", func(t *testing.T) {",
				},
				{
					LineNumber: targetLine,
					Content:    "\t\tbt := getBacktrace(0)",
				},
				{
					LineNumber: targetLine + 1,
					Content:    "\t\tif len(bt) == 0 {",
				},
			},
		}

		if (expected.File != bt[0].File) || (expected.Line != bt[0].Line) || (expected.Function != bt[0].Function) {
			t.Fatalf("wrong backtrace:\n\texpected: %+v\n\tactual: %+v", expected, bt[0])
		}
		if len(expected.SourceCode) != len(bt[0].SourceCode) {
			t.Fatalf("wrong source code length:\n\texpected: %d\n\tactual: %d", len(expected.SourceCode), len(bt[0].SourceCode))
		} else if !reflect.DeepEqual(expected.SourceCode[0], bt[0].SourceCode[0]) || !reflect.DeepEqual(expected.SourceCode[1], bt[0].SourceCode[1]) || !reflect.DeepEqual(expected.SourceCode[2], bt[0].SourceCode[2]) {
			t.Fatalf("wrong source code:\n\texpected: %+v\n\tactual: %+v", expected.SourceCode, bt[0].SourceCode)
		}
	})

	t.Run("race test", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 1; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				bt := getBacktrace(0)
				if len(bt) == 0 {
					t.Fatalf("received empty backtrace")
				}
			}(i)
		}
		wg.Wait()
	})
}
