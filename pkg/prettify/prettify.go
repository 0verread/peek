// custom package to make json response colorful
package prettify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

const (
	startMap   = "{\n"
	endMap     = "\n}"
	startArray = "[\n"
	endArray   = "\n]"
	emptyArray = startArray + endArray
)

type VerbFunc func(a ...interface{}) string
type UrlFunc func(a ...interface{}) string

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) colorize(c color.Attribute, s string) string {
	return color.New(c).SprintFunc()(s)
}

func Verb(verb string) VerbFunc {
	return color.New(VerbColor).SprintFunc()
}

func Url(url string) UrlFunc {
	return color.New(UrlColor).SprintFunc()
}

func (f *Formatter) colorString(str string, buf *bytes.Buffer) {
	strBytes, _ := json.Marshal(str)
	buf.WriteString(f.colorize(color.FgGreen, string(strBytes)))
}

func (f *Formatter) colorArray(arr []interface{}, buf *bytes.Buffer) error {
	buf.WriteString(startArray)
	for i, v := range arr {
		if i > 0 {
			buf.WriteString(",")
		}
		if err := f.marshalValue(v, buf); err != nil {
			return err
		}
	}
	buf.WriteString(endArray)
	return nil
}

func (f *Formatter) colorMap(m map[string]interface{}, buf *bytes.Buffer) error {
	buf.WriteString(startMap)
	var first = true
	for k, v := range m {
		if !first {
			buf.WriteString(", \n")
		}
		first = false
		// color the key
		buf.WriteString(f.colorize(color.FgWhite, fmt.Sprintf("  %q: ", k)))
		if err := f.marshalValue(v, buf); err != nil {
			return err
		}
	}
	buf.WriteString(endMap)
	return nil
}

func (f *Formatter) colorArrayMap(arr []map[string]interface{}, buf *bytes.Buffer) error {
	buf.WriteString(startArray)
	for i, v := range arr {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("\n  ")
		if err := f.colorMap(v, buf); err != nil {
			return err
		}
	}
	buf.WriteString(endArray)
	return nil
}

func (f *Formatter) marshalValue(value interface{}, buf *bytes.Buffer) error {
	switch v := value.(type) {
	case string:
		f.colorString(v, buf)
	case float64:
		buf.WriteString(f.colorize(color.FgCyan, strconv.FormatFloat(v, 'f', -1, 64)))
	case bool:
		buf.WriteString(f.colorize(color.FgYellow, strconv.FormatBool(v)))
	case nil:
		buf.WriteString(f.colorize(color.FgMagenta, "null"))
	case []interface{}:
		return f.colorArray(v, buf)
	case map[string]interface{}:
		return f.colorMap(v, buf)
	case []map[string]interface{}:
		return f.colorArrayMap(v, buf)
	case json.Number:
		buf.WriteString(f.colorize(color.FgCyan, v.String()))
	default:
		return fmt.Errorf("Type is not supported")
	}
	return nil
}

func (f *Formatter) Prettify(jsonObj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := f.marshalValue(jsonObj, &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Prettify(jsonObj interface{}) ([]byte, error) {
	return NewFormatter().Prettify(jsonObj)
}
