package validator

import "testing"

func TestValidateRequiredMissingValue(t *testing.T) {
	env := map[string]string{}

	schema := Schema{
		"DATABASE_URL": {
			Required: true,
		},
	}

	results := Validate(env, schema)

	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}

	if results[0].Level != LevelFail {
		t.Fatalf("Expected LevelFail, got %s", results[0].Level)
	}
}

func TestValidateAllowedValues(t *testing.T) {
	env := map[string]string{
		"NODE_ENV": "banana",
	}

	schema := Schema{
		"NODE_ENV": {
			AllowedValues: []string{"development", "test", "production"},
		},
	}

	results := Validate(env, schema)

	if results[0].Level != LevelFail {
		t.Fatalf("Expected LevelFail, got %s", results[0].Level)
	}
}

func TestValidateNumberType(t *testing.T) {
	env := map[string]string{
		"SMTP_PORT": "not-a-number",
	}

	schema := Schema{
		"SMTP_PORT": {
			Type: "number",
		},
	}

	results := Validate(env, schema)

	if results[0].Level != LevelFail {
		t.Fatalf("Expected LevelFail, got %s", results[0].Level)
	}
}

func TestValidateValidValue(t *testing.T) {
	env := map[string]string{
		"SMTP_PORT": "587",
	}

	schema := Schema{
		"SMTP_PORT": {
			Type: "number",
		},
	}

	results := Validate(env, schema)

	if results[0].Level != LevelOK {
		t.Fatalf("Expected LevelOK, got %s", results[0].Level)
	}
}
