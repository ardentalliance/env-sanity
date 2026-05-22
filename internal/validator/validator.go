package validator

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// defines rules, schema, levels
type Rule struct {
	Required      bool     `json:"required"`
	MinLength     int      `json:"minLength"`
	AllowedValues []string `json:"allowedValues"`
	Type          string   `json:"type"`
}

type Schema map[string]Rule

type Level string

const (
	LevelOK   Level = "OK"
	LevelWarn Level = "WARN"
	LevelFail Level = "FAIL"
)

type Result struct {
	Level   Level
	Key     string
	Message string
}

// load schema file
func LoadSchema(path string) (Schema, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not open schema file %q: %w", path, err)
	}

	var schema Schema
	if err := json.Unmarshal(content, &schema); err != nil {
		return nil, fmt.Errorf("could not parse schema file %q: %w", path, err)
	}

	return schema, nil
}

// validate env file based on values set in schema
func Validate(env map[string]string, schema Schema) []Result {
	results := make([]Result, 0, len(schema))

	for key, rule := range schema {
		value, exists := env[key]

		// required value, not set
		if rule.Required && (!exists || value == "") {
			results = append(results, Result{
				Level:   LevelFail,
				Key:     key,
				Message: "is required but missing",
			})
			continue
		}

		// optional value, not set
		if !exists || value == "" {
			results = append(results, Result{
				Level:   LevelWarn,
				Key:     key,
				Message: "is optional and not set",
			})
			continue
		}

		// value too short
		if rule.MinLength > 0 && len(value) < rule.MinLength {
			results = append(results, Result{
				Level:   LevelFail,
				Key:     key,
				Message: fmt.Sprintf("must be at least %d characters", rule.MinLength),
			})
			continue
		}

		// not one of the allowed values
		if len(rule.AllowedValues) > 0 && !contains(rule.AllowedValues, value) {
			results = append(results, Result{
				Level:   LevelFail,
				Key:     key,
				Message: fmt.Sprintf("value must be one of : %s", strings.Join(rule.AllowedValues, ", ")),
			})
			continue
		}

		// wrong value type
		if rule.Type != "" {
			if err := validateType(value, rule.Type); err != nil {
				results = append(results, Result{
					Level:   LevelFail,
					Key:     key,
					Message: err.Error(),
				})
				continue
			}
		}

		results = append(results, Result{
			Level:   LevelOK,
			Key:     key,
			Message: "is valid and set",
		})
	}

	return results
}

// if any checks fail
func HasFailures(results []Result) bool {
	for _, result := range results {
		if result.Level == LevelFail {
			return true
		}
	}

	return false
}

// type validation
func validateType(value string, expectedType string) error {
	switch expectedType {
	case "string":
		return nil
	case "number":
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return fmt.Errorf("value must be a number: %w", err)
		}
	case "boolean":
		normalized := strings.ToLower(value)
		if normalized != "true" && normalized != "false" && normalized != "1" && normalized != "0" {
			return fmt.Errorf("value must be a boolean")
		}
	default:
		return fmt.Errorf("uses unsupported type %q", expectedType)
	}

	return nil
}

// value contains x
func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}

	return false
}
