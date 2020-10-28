package hyperdeck

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Parse(payload []byte, res interface{}) error {
	buff := bytes.NewBuffer(payload)
	_, err := buff.ReadString('\n')
	// Ignore first line (header)
	if err != nil {
		return errors.Wrap(err, "fail to read first line")
	}

	temp := make(map[string]string)

	// Step 1: Transform the message to a map
	for {
		line, err := buff.ReadString('\n')
		if err != nil {
			return errors.Wrap(err, "fail to read line")
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		fields := strings.SplitN(line, ":", 2)
		if len(fields) != 2 {
			return fmt.Errorf("invalid line: %s", line)
		}
		temp[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
	}

	// Step 2: Fill the different fields

	elems := reflect.ValueOf(res).Elem()
	types := elems.Type()
	for i := 0; i < types.NumField(); i++ {
		fieldType := types.Field(i)
		field := elems.FieldByName(fieldType.Name)
		lookup := fieldType.Tag.Get("header")
		if lookup == "" {
			continue
		}
		value, ok := temp[lookup]
		if !ok {
			continue
		}
		if fieldType.Type.Name() == "string" {
			field.SetString(value)
		} else if fieldType.Type.Name() == "int" {
			if value != "none" {
				val, err := strconv.Atoi(value)
				if err != nil {
					return errors.Wrapf(err, "invalid value for %s", fieldType.Name)
				}
				field.SetInt(int64(val))
			}
		} else {
			parser, ok := field.Addr().Interface().(StringParser)
			if ok {
				err := parser.FromString(value)
				if err != nil {
					return errors.Wrapf(err, "invalid value for %s", fieldType.Name)
				}
			}
		}
	}
	return nil
}
