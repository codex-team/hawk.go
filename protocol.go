package hawk

// ErrorReport is a report about an error that is sent to Hawk.
// easyjson:json
type ErrorReport struct {
	// Token is the Hawk access token.
	Token string `json:"token"`
	// CatcherType is the type of this Catcher.
	CatcherType string `json:"catcherType"`
	// Payload is the information about the error.
	Payload Payload `json:"payload"`
}

// Payload is the information about the error.
type Payload struct {
	// Title is the error name.
	Title string `json:"title"`
	// Timestamp represents time when the error was caught by the Catcher.
	Timestamp string `json:"timestamp"`
	// Severity is the error's severity level.
	Severity int `json:"level,omitempty"`
	// Backtrace contains information about the function calls that caused the
	// error.
	Backtrace []Backtrace `json:"backtrace"`
}

// Backtrace contains information about the function calls that caused the
// error.
type Backtrace struct {
	// File is the path to file where the error was caught.
	File string `json:"file"`
	// Line is the number of line which caused the error.
	Line int `json:"line"`
	// Function is the name of function where error was caught.
	Function string `json:"function"`
	// SourceCode contains the line which caused the error, the previous and the
	// next lines.
	SourceCode []SourceCode `json:"sourceCode,omitempty"`
}

// SourceCode contains the line which caused the error, the previous and the
// next lines.
type SourceCode struct {
	// LineNumber is the number of line which caused the error.
	LineNumber int `json:"line"`
	// Content is the line itself.
	Content string `json:"content"`
}
