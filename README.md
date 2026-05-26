# env-sanity

`env-sanity` is a small Go CLI for validating `.env` files against a JSON schema.

It checks for missing required variables, invalid values, and simple type mismatches.

<!-- TOC -->
* [env-sanity](#env-sanity)
  * [Features](#features)
  * [Installation](#installation)
    * [Option 1: Build from source code](#option-1-build-from-source-code)
      * [Clone the repository:](#clone-the-repository)
      * [Build the binary](#build-the-binary)
      * [Run the software](#run-the-software)
    * [Option 2: Install with Go](#option-2-install-with-go)
    * [Option 3: Download a release binary](#option-3-download-a-release-binary)
  * [Example Usage](#example-usage)
    * [Use a different .env file](#use-a-different-env-file)
    * [Use a different schema file](#use-a-different-schema-file)
    * [Print results as JSON](#print-results-as-json)
  * [Example Output](#example-output)
    * [Standard Output](#standard-output)
    * [JSON Output](#json-output)
  * [Quick-Start Guide](#quick-start-guide)
  * [Example Schema Format](#example-schema-format)
  * [Supported Schema Role Properties](#supported-schema-role-properties)
  * [Development](#development)
    * [Run the CLI locally without building the binaries](#run-the-cli-locally-without-building-the-binaries)
    * [Run tests](#run-tests)
    * [Build locally](#build-locally)
  * [Roadmap](#roadmap)
  * [Ethical use](#ethical-use)
<!-- TOC -->

---

## Features

- Validate required environment variables
- Check minimum length
- Check allowed values
- Check simple types: `string`, `number`, `boolean`
- Warn about likely placeholder values in sensitive variables
- Allow placeholder warnings to be disabled per variable
- Returns a non-zero exit code when validation fails

---

## Installation

### Option 1: Build from source code

#### Clone the repository:
```bash
git clone https://github.com/ardentalliance/env-sanity.git
```

#### Build the binary
```bash
go build -o env-sanity ./cmd/env-sanity
```

#### Run the software
```bash
./env-sanity check
```

### Option 2: Install with Go

If you have Go installed on your machine, you can directly install the latest version of the CLI with:
```bash
go install github.com/ardentalliance/env-sanity/cmd/env-sanity@latest
```

After installation, run:
```bash
env-sanity check
```

This only works if your Go binary directory is in your PATH.

### Option 3: Download a release binary

Download the release binary for your OS and architecture from the [GitHub Releases Page](https://github.com/ardentalliance/env-sanity/releases).

Then run the tool with:
```bash
env-sanity check
```

---

## Example Usage

`env-sanity` checks a `.env` file against an `env.schema.json` file. 

By default, it expects both files to be in the current folder:

```text
project-root
├── .env
├── env.schema.json
└── ...
```

From your project root, run:
```bash
env-sanity check
```
This checks `.env` against `env.schema.json`

### Use a different .env file

```bash
env-sanity check -env .env.file.location
```

### Use a different schema file

```bash
env-sanity check -schema config/env.schema.json
```

### Print results as JSON

```bash
env-sanity check -json
```

## Example Output

### Standard Output
```bash
OK    DATABASE_URL             is set
FAIL  SESSION_SECRET           must be at least 32 characters
WARN  NODE_ENV                 is optional and not set
OK    SMTP_PORT                is set
WARN  FEATURE_FLAG_EMAIL       looks like a placeholder secret       
```
### JSON Output
```json
[
  {
    "level": "OK",
    "key": "DATABASE_URL",
    "message": "is valid and set"
  },
  {
    "level": "FAIL",
    "key": "SESSION_SECRET",
    "message": "must be at least 32 characters"
  },
  {
    "level": "WARN",
    "key": "NODE_ENV",
    "message": "looks like a placeholder secret"
  },
  {
    "level": "WARN",
    "key": "SMTP_PORT",
    "message": "is optional and not set"
  }
]
```


---

## Quick-Start Guide

After installing the software, the basic workflow for using this software is:

1. Copy the `.env.example` file from `examples/` into your project root.
2. Replace the values there with the environment variables you'll be using for your project 
3. Copy the `env.schema.json` file from `examples` into your project root.
4. Replace the values there to describe the validation needed for your variables
5. Run `env-sanity check` from the command line.
6. Fix anything marked as `FAIL`.

---

## Example Schema Format

`env-sanity` validates variables based on a JSON schema file.

Example `env.schema.json`:

```json
{
    "DATABASE_URL": {
        "required": true
    },
    "SESSION_SECRET": {
        "required": true,
        "minLength": 32,
        "allowPlaceholder": true
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

## Supported Schema Role Properties

| Property           | Type         | Description                                                                            | 
|--------------------|--------------|----------------------------------------------------------------------------------------|
| `required`         | boolean      | Marks the variable as required. If the variable is missing or empty, validation fails. |
| `minLength`        | number       | Requires the variable to have at least the specified number of characters.             | 
| `allowedValues`    | string array | Variable must have one of the listed values to pass validation                         |
| `type`             | string       | Checks simple value types. Supported values: string, number, boolean                   |
| `allowPlaceholder` | boolean      | Disables or allows placeholder-secret warnings for this variable                       |

---

## Development

Instructions for working with the source code:

### Run the CLI locally without building the binaries

```bash
go run ./cmd/env-sanity check
```

### Run tests

```bash
go test ./...
```

### Build locally

```bash
go build -o env-sanity ./cmd/env-sanity
```

---

## Roadmap


1. Add .env.example generation
2. Add GitHub Actions test workflow

---

## Ethical use

This project is provided as open-source software under the MIT License.
I ask users not to use this software for harm, surveillance, human rights abuses,
or military aggression.