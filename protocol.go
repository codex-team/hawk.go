package hawk

import "github.com/mailru/easyjson"

const DefaultType = "error"
const ManualType = "manual"

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

type ErrorContext easyjson.RawMessage

type AffectedUser struct {
	// Internal user's identifier inside an app
	Id string `json:"id"`
	// User public name
	Name string `json:"name"`
	// URL for user's details page
	URL string `json:"url"`
	// User's public picture
	Image string `json:"image"`
}

func (au *AffectedUser) isEmpty() bool {
	return au.Id == "" && au.Name == "" && au.URL == "" && au.Image == ""
}

// Payload is the information about the error.
type Payload struct {
	// Title is the error name.
	Title string `json:"title"`
	// Event type (severity level)
	Type string `json:"type"`
	// Backtrace contains information about the function calls that caused the error.
	Backtrace []Backtrace `json:"backtrace"`
	// Current release (aka version, revision) of an application
	Release string `json:"release"`
	// Current authenticated user
	User AffectedUser `json:"user"`
	// Any other information collected and passed by user
	Context easyjson.RawMessage `json:"context"`
	// Catcher version
	CatcherVersion string `json:"catcherVersion"`
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
