package operations

import (
	"bytes"
	"fmt"
	"github.com/sqlbunny/sqlschema/schema"
)

func sqlName(schema, name string) string {
	if schema == "" {
		return fmt.Sprintf("\"%s\"", name)
	}
	return fmt.Sprintf("\"%s\".\"%s\"", schema, name)
}

func esc(s string) string {
	return fmt.Sprintf("%#v", s)
}

func dumpBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func columnList(columns []string) string {
	var buf bytes.Buffer
	for i, c := range columns {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("\"")
		buf.WriteString(c)
		buf.WriteString("\"")
	}
	return buf.String()
}

func indexType(t schema.IndexType) string {
	switch t {
	case schema.DefaultIndex:
		return ""
	case schema.GISTIndex:
		return "USING GIST"
	default:
		return ""
	}
}
