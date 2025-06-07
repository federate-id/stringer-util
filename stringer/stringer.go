package stringer

import (
	"fmt"
	"reflect"
	"strings"
)

func ToStringWithTags(obj any) string {
	return toStringWithTags(obj, 0)
}

func toStringWithTags(obj any, depth int) string {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Sprintf("%v", v)
	}

	var b strings.Builder
	b.WriteString(t.Name())
	b.WriteString("{")
	first := true

	for i := range t.NumField() {
		field := t.Field(i)
		value := v.Field(i)

		if field.PkgPath != "" {
			// this is a private field which cannot be reflected with value.Interface()
			continue
		}

		tag := field.Tag.Get("stringer")
		if tag == "" {
			// only annotated fields are included
			continue
		}

		tagParts := parseTag(tag)
		if !first {
			b.WriteString(", ")
		}
		first = false

		fieldName := tagParts["name"]
		if fieldName == "" {
			fieldName = field.Name
		}

		b.WriteString(fieldName)
		b.WriteString(": ")

		switch tagParts["mode"] {
		case "include":
			b.WriteString(fmt.Sprintf("%v", value.Interface()))
		case "masked":
			if tagParts["length"] == "true" && value.Kind() == reflect.String {
				b.WriteString(fmt.Sprintf("***** (len=%d)", len(value.String())))
			} else {
				if value.Kind() == reflect.String {
					if len(value.String()) > 0 {
						b.WriteString("*****")
					} else {
						b.WriteString("\"\"")
					}
				} else {
					b.WriteString(fmt.Sprintf("%v", value.Interface()))
				}

			}
		case "type":
			b.WriteString(fmt.Sprintf("<%s>", field.Type.String()))
		case "nested":
			b.WriteString(toStringWithTags(value.Interface(), depth+1))
		default:
			b.WriteString("\"\"")
		}
	}
	b.WriteString("}")
	return b.String()
}

func parseTag(tag string) map[string]string {
	parts := strings.Split(tag, ",")
	res := make(map[string]string)
	for _, part := range parts {
		if part == "include" || part == "masked" || part == "type" || part == "nested" {
			res["mode"] = part
		} else if part == "length" {
			res["length"] = "true"
		} else if strings.Contains(part, "=") {
			kv := strings.SplitN(part, "=", 2)
			res[kv[0]] = kv[1]
		}
	}
	return res
}
