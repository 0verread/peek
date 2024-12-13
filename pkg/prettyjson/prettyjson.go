package prettyjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

const (
	startMap   = "{"
	endMap     = "}"
	startArray = "["
	endArray   = "]"
	emptyArray = startArray + endArray
)

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) colorize(c color.Attribute, s string) string {
	return color.New(c).SprintFunc()(s)
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
			buf.WriteString(",")
		}
		first = false
		// color the key
		buf.WriteString(f.colorize(color.FgWhite, fmt.Sprintf("%q: ", k)))
		if err := f.marshalValue(v, buf); err != nil {
			return err
		}
	}
	buf.WriteString(endMap)
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
	case json.Number:
		buf.WriteString(f.colorize(color.FgCyan, v.String()))
	default:
		return fmt.Errorf("Type is not supported")
	}
	return nil
}

func (f *Formatter) marshalArray(arr []interface{}, buf *bytes.Buffer, depth int) error {
	if len(arr) == 0 {
		buf.WriteString(emptyArray)
	}
	buf.WriteString(startArray)

	for i, v := range arr {
		if i > 0 {
			buf.WriteByte(',')
		}
		if err := f.marshalValue(v, buf); err != nil {
			return err
		}
	}
	buf.WriteString(endArray)
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
	fmt.Println(jsonObj)
	return NewFormatter().Prettify(jsonObj)
}
