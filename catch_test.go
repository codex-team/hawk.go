package hawk

import (
	"os"
	"reflect"
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
		bt[0].SourceCode, err = readSourceCode(file, bt[0].Line)
		if err != nil {
			t.Fatalf("failed to read source code: %s", err.Error())
		}

		pwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("failed to get pwd: %s", err.Error())
		}

		targetLine := 12
		expected := Backtrace{
			File: pwd + "/catch_test.go",
			Line: targetLine,
			SourceCode: [SourceCodeLines]SourceCode{
				{
					LineNumber: targetLine - 1,
					Content:    "t.Run(\"simple case\", func(t *testing.T) {",
				},
				{
					LineNumber: targetLine,
					Content:    "bt := getBacktrace(0)",
				},
				{
					LineNumber: targetLine + 1,
					Content:    "if len(bt) == 0 {",
				},
			},
		}

		if !reflect.DeepEqual(expected, bt[0]) {
			t.Fatalf("wrong backtrace:\n\texpected: %+v\n\tactual: %+v", expected, bt[0])
		}
	})
}
