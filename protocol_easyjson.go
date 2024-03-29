// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package hawk

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE4425964DecodeGithubComCodexTeamHawkGo(in *jlexer.Lexer, out *SourceCode) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "line":
			out.LineNumber = int(in.Int())
		case "content":
			out.Content = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE4425964EncodeGithubComCodexTeamHawkGo(out *jwriter.Writer, in SourceCode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"line\":"
		out.RawString(prefix[1:])
		out.Int(int(in.LineNumber))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SourceCode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SourceCode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SourceCode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SourceCode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo(l, v)
}
func easyjsonE4425964DecodeGithubComCodexTeamHawkGo1(in *jlexer.Lexer, out *Payload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			out.Title = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "backtrace":
			if in.IsNull() {
				in.Skip()
				out.Backtrace = nil
			} else {
				in.Delim('[')
				if out.Backtrace == nil {
					if !in.IsDelim(']') {
						out.Backtrace = make([]Backtrace, 0, 1)
					} else {
						out.Backtrace = []Backtrace{}
					}
				} else {
					out.Backtrace = (out.Backtrace)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Backtrace
					(v1).UnmarshalEasyJSON(in)
					out.Backtrace = append(out.Backtrace, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "release":
			out.Release = string(in.String())
		case "user":
			(out.User).UnmarshalEasyJSON(in)
		case "context":
			(out.Context).UnmarshalEasyJSON(in)
		case "catcherVersion":
			out.CatcherVersion = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE4425964EncodeGithubComCodexTeamHawkGo1(out *jwriter.Writer, in Payload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"backtrace\":"
		out.RawString(prefix)
		if in.Backtrace == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Backtrace {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"release\":"
		out.RawString(prefix)
		out.String(string(in.Release))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		(in.User).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"context\":"
		out.RawString(prefix)
		(in.Context).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"catcherVersion\":"
		out.RawString(prefix)
		out.String(string(in.CatcherVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Payload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Payload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Payload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Payload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo1(l, v)
}
func easyjsonE4425964DecodeGithubComCodexTeamHawkGo2(in *jlexer.Lexer, out *ErrorReport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		case "catcherType":
			out.CatcherType = string(in.String())
		case "payload":
			(out.Payload).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE4425964EncodeGithubComCodexTeamHawkGo2(out *jwriter.Writer, in ErrorReport) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	{
		const prefix string = ",\"catcherType\":"
		out.RawString(prefix)
		out.String(string(in.CatcherType))
	}
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix)
		(in.Payload).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorReport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorReport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorReport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorReport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo2(l, v)
}
func easyjsonE4425964DecodeGithubComCodexTeamHawkGo3(in *jlexer.Lexer, out *Backtrace) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "file":
			out.File = string(in.String())
		case "line":
			out.Line = int(in.Int())
		case "function":
			out.Function = string(in.String())
		case "sourceCode":
			if in.IsNull() {
				in.Skip()
				out.SourceCode = nil
			} else {
				in.Delim('[')
				if out.SourceCode == nil {
					if !in.IsDelim(']') {
						out.SourceCode = make([]SourceCode, 0, 2)
					} else {
						out.SourceCode = []SourceCode{}
					}
				} else {
					out.SourceCode = (out.SourceCode)[:0]
				}
				for !in.IsDelim(']') {
					var v4 SourceCode
					(v4).UnmarshalEasyJSON(in)
					out.SourceCode = append(out.SourceCode, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE4425964EncodeGithubComCodexTeamHawkGo3(out *jwriter.Writer, in Backtrace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"file\":"
		out.RawString(prefix[1:])
		out.String(string(in.File))
	}
	{
		const prefix string = ",\"line\":"
		out.RawString(prefix)
		out.Int(int(in.Line))
	}
	{
		const prefix string = ",\"function\":"
		out.RawString(prefix)
		out.String(string(in.Function))
	}
	if len(in.SourceCode) != 0 {
		const prefix string = ",\"sourceCode\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.SourceCode {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Backtrace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Backtrace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Backtrace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Backtrace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo3(l, v)
}
func easyjsonE4425964DecodeGithubComCodexTeamHawkGo4(in *jlexer.Lexer, out *AffectedUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "image":
			out.Image = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE4425964EncodeGithubComCodexTeamHawkGo4(out *jwriter.Writer, in AffectedUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AffectedUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AffectedUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4425964EncodeGithubComCodexTeamHawkGo4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AffectedUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AffectedUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4425964DecodeGithubComCodexTeamHawkGo4(l, v)
}
