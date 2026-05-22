# env-sanity

`env-sanity` is a small Go CLI for validating `.env` files against a JSON schema.

It checks for missing required variables, invalid values, and simple type mismatches.

---

## Features

- Validate required environment variables
- Check minimum length
- Check allowed values
- Check simple types: `string`, `number`, `boolean`
- Returns a non-zero exit code when validation fails

---

## Example Usage

```bash
env-sanity check
```

### Example Output:

```bash
OK    DATABASE_URL             is set
FAIL  SESSION_SECRET           must be at least 32 characters
WARN  NODE_ENV                 is optional and not set
OK    SMTP_PORT                is set
```

---

## Example Schema Format

```json
{
    "DATABASE_URL": {
        "required": true
    },
    "SESSION_SECRET": {
        "required": true,
        "minLength": 32
    },
    "NODE_ENV": {
        "required": false,
        "allowedValues": ["development", "test", "production"]
    },
    "SMTP_PORT": {
        "required": false,
        "type": "number"
    }
}
```

---

## Development

### Run locally:

```bash
go run ./cmd/env-sanity check
```

### Run tests:

```bash
go test ./...
```


### Build

```bash
go build -o env-sanity ./cmd/env-sanity
```

---

## Roadmap

1. Add JSON Output
2. Add .env.example generation
3. Warn about weak placeholder secrets
4. Add GitHub Actions test workflow

---

## Ethical use

This project is provided as open-source software under the MIT License.
I ask users not to use this software for harm, surveillance, human rights abuses,
or military aggression.