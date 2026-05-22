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

---

## License


Copyright 2026 @ardentalliance

Permission is hereby granted, free of charge, to any person obtaining a copy of this software
and associated documentation files (the “Software”), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software
is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE
AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, 
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.


