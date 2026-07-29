// Package convert translates SARIF reports between the versions supported by
// go-sarif.
//
// The 2.1.0 and 2.2 object models are almost identical: every shared field uses
// the same JSON name, so conversion is driven by a JSON round trip rather than
// several hundred hand-written field copies. Fields that the target version has
// no home for are pruned, and every pruned field is reported as a Loss.
//
// Upgrading (2.1.0 to 2.2) is lossless. Downgrading (2.2 to 2.1.0) is not: 2.2
// added a report-level guid and relatedLocations on notification, neither of
// which 2.1.0 can represent. By default those fields are dropped and reported;
// pass WithStrictConversion to turn any loss into an error instead.
//
// Note that a report built by v22.NewReport always carries a generated guid, so
// a strict downgrade of such a report will report that guid as a loss. Clear
// Report.Guid first if that is not wanted.
package convert

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	v210 "github.com/owenrumney/go-sarif/v3/pkg/report/v210/sarif"
	v22 "github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
)

// Loss records a single field that could not be represented in the target
// version and was dropped during conversion.
type Loss struct {
	// Path is the JSON path of the dropped field, for example
	// "runs[0].invocations[0].toolExecutionNotifications[0].relatedLocations".
	Path string

	// Value is the dropped value, so callers can preserve it out of band.
	Value any
}

func (l Loss) String() string {
	return fmt.Sprintf("%s has no equivalent in the target version", l.Path)
}

// LossyConversionError is returned by a strict conversion that had to drop one
// or more fields.
type LossyConversionError struct {
	Losses []Loss
}

func (e *LossyConversionError) Error() string {
	paths := make([]string, 0, len(e.Losses))
	for _, l := range e.Losses {
		paths = append(paths, l.Path)
	}
	return fmt.Sprintf("lossy conversion: %s have no equivalent in the target version", strings.Join(paths, ", "))
}

// Option configures a conversion.
type Option func(*config)

type config struct {
	strict bool
	onLoss func(Loss)
}

// WithStrictConversion causes a conversion that would drop data to fail with a
// *LossyConversionError instead of dropping it.
func WithStrictConversion() Option {
	return func(cfg *config) {
		cfg.strict = true
	}
}

// WithLossHandler calls handler once for each field dropped during a non-strict
// conversion, so callers can log or preserve what could not be represented.
func WithLossHandler(handler func(Loss)) Option {
	return func(cfg *config) {
		cfg.onLoss = handler
	}
}

const (
	v210Schema = "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/main/sarif-2.1/schema/sarif-schema-2.1.0.json"
	v22Schema  = "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/main/sarif-2.2/schema/sarif-2-2.schema.json"

	// defaultLanguage is the default the SARIF spec defines for language.
	defaultLanguage = "en-US"
)

// ToV22 converts a 2.1.0 report to 2.2. The conversion is lossless, so
// WithStrictConversion never causes it to fail.
//
// A report-level guid is generated, since 2.2 requires one and 2.1.0 has none
// to carry over.
func ToV22(report *v210.Report, options ...Option) (*v22.Report, error) {
	if report == nil {
		return nil, errors.New("cannot convert a nil report")
	}

	pruned, err := transcode(report, reflect.TypeOf(v22.Report{}), options...)
	if err != nil {
		return nil, err
	}

	converted, err := v22.FromBytes(pruned)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the converted report: %w", err)
	}

	converted.Version = "2.2"
	converted.Schema = v22Schema
	if converted.Guid == "" {
		converted.Guid = v22.NewGuid()
	}

	return converted, nil
}

// ToV210 converts a 2.2 report to 2.1.0. 2.2 fields that 2.1.0 cannot represent
// are dropped, and reported to the handler given to WithLossHandler. With
// WithStrictConversion the conversion fails with a *LossyConversionError
// instead of dropping anything.
func ToV210(report *v22.Report, options ...Option) (*v210.Report, error) {
	if report == nil {
		return nil, errors.New("cannot convert a nil report")
	}

	pruned, err := transcode(report, reflect.TypeOf(v210.Report{}), options...)
	if err != nil {
		return nil, err
	}

	converted, err := v210.FromBytes(pruned)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the converted report: %w", err)
	}

	converted.Version = "2.1.0"
	converted.Schema = v210Schema
	defaultLanguages(reflect.ValueOf(converted))

	return converted, nil
}

// defaultLanguages fills in the SARIF default language for every language field
// left empty by the conversion. 2.2 made language optional, but the 2.1.0 model
// always serializes it and the 2.1.0 schema constrains it to a culture code, so
// an empty value would produce an invalid report. Run and ToolComponent both
// carry one, and ToolComponent appears throughout a run, so this walks the whole
// report rather than naming each site.
func defaultLanguages(value reflect.Value) {
	switch value.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !value.IsNil() {
			defaultLanguages(value.Elem())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			defaultLanguages(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			defaultLanguages(value.MapIndex(key))
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			if value.Type().Field(i).Name == "Language" && field.Kind() == reflect.String {
				if field.CanSet() && field.String() == "" {
					field.SetString(defaultLanguage)
				}
				continue
			}
			defaultLanguages(field)
		}
	}
}

// transcode marshals source and prunes any field that target cannot represent,
// returning the resulting JSON.
func transcode(source any, target reflect.Type, options ...Option) ([]byte, error) {
	cfg := &config{}
	for _, opt := range options {
		opt(cfg)
	}

	marshaled, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("failed to encode the source report: %w", err)
	}

	// UseNumber keeps numeric literals exact; decoding into any would otherwise
	// route every number through float64 and reformat large values.
	var tree any
	decoder := json.NewDecoder(bytes.NewReader(marshaled))
	decoder.UseNumber()
	if err := decoder.Decode(&tree); err != nil {
		return nil, fmt.Errorf("failed to decode the source report: %w", err)
	}

	var losses []Loss
	tree = prune(tree, target, "", &losses)

	if cfg.strict && len(losses) > 0 {
		return nil, &LossyConversionError{Losses: losses}
	}
	if cfg.onLoss != nil {
		for _, loss := range losses {
			cfg.onLoss(loss)
		}
	}

	pruned, err := json.Marshal(tree)
	if err != nil {
		return nil, fmt.Errorf("failed to re-encode the converted report: %w", err)
	}

	return pruned, nil
}

// prune walks a decoded JSON tree alongside the Go type it must decode into,
// removing any object key with no corresponding field and recording it as a
// loss. It mirrors how encoding/json matches keys to fields.
func prune(node any, target reflect.Type, path string, losses *[]Loss) any {
	for target != nil && target.Kind() == reflect.Pointer {
		target = target.Elem()
	}

	// An any-typed destination accepts whatever it is given, so stop here.
	if target == nil || target.Kind() == reflect.Interface {
		return node
	}

	switch typed := node.(type) {
	case map[string]any:
		switch target.Kind() {
		case reflect.Struct:
			fields := jsonFields(target)
			for key, value := range typed {
				fieldType, ok := fields[key]
				if !ok {
					// Dropping an absent or empty value loses nothing, and the two
					// versions disagree about which fields carry omitempty, so only
					// a populated value counts as a loss.
					if !isEmpty(value) {
						*losses = append(*losses, Loss{Path: join(path, key), Value: value})
					}
					delete(typed, key)
					continue
				}
				typed[key] = prune(value, fieldType, join(path, key), losses)
			}
		case reflect.Map:
			for key, value := range typed {
				typed[key] = prune(value, target.Elem(), join(path, key), losses)
			}
		}
		return typed

	case []any:
		if target.Kind() != reflect.Slice && target.Kind() != reflect.Array {
			return typed
		}
		for i, value := range typed {
			typed[i] = prune(value, target.Elem(), fmt.Sprintf("%s[%d]", path, i), losses)
		}
		return typed
	}

	return node
}

// isEmpty reports whether a decoded JSON value carries no information, and so
// can be dropped without losing anything.
func isEmpty(value any) bool {
	switch typed := value.(type) {
	case nil:
		return true
	case string:
		return typed == ""
	case []any:
		return len(typed) == 0
	case map[string]any:
		return len(typed) == 0
	}
	return false
}

// jsonFields maps the JSON name of each exported field of a struct to its type.
func jsonFields(target reflect.Type) map[string]reflect.Type {
	fields := make(map[string]reflect.Type, target.NumField())
	for i := 0; i < target.NumField(); i++ {
		field := target.Field(i)
		if field.PkgPath != "" {
			continue // unexported
		}

		name := field.Name
		if tag := field.Tag.Get("json"); tag != "" {
			tagName, _, _ := strings.Cut(tag, ",")
			if tagName == "-" {
				continue
			}
			if tagName != "" {
				name = tagName
			}
		}
		fields[name] = field.Type
	}
	return fields
}

func join(path, key string) string {
	if path == "" {
		return key
	}
	return path + "." + key
}
