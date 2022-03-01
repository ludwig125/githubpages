// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package args

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

func easyjson1d13fa9bDecodeGithubComLudwig125GithubpagesDocsCalc3TinygoTrickyArgsArgs(in *jlexer.Lexer, out *Args) {
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
		case "x":
			out.X = string(in.String())
		case "y":
			out.Y = string(in.String())
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
func easyjson1d13fa9bEncodeGithubComLudwig125GithubpagesDocsCalc3TinygoTrickyArgsArgs(out *jwriter.Writer, in Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x\":"
		out.RawString(prefix[1:])
		out.String(string(in.X))
	}
	{
		const prefix string = ",\"y\":"
		out.RawString(prefix)
		out.String(string(in.Y))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1d13fa9bEncodeGithubComLudwig125GithubpagesDocsCalc3TinygoTrickyArgsArgs(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1d13fa9bEncodeGithubComLudwig125GithubpagesDocsCalc3TinygoTrickyArgsArgs(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1d13fa9bDecodeGithubComLudwig125GithubpagesDocsCalc3TinygoTrickyArgsArgs(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1d13fa9bDecodeGithubComLudwig125GithubpagesDocsCalc3TinygoTrickyArgsArgs(l, v)
}
