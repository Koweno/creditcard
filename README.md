# CreditCard Tool

A command-line tool for validating, generating, and managing credit card numbers. This tool provides functionalities such as validating credit card numbers, generating possible card numbers, retrieving card information, and issuing new card numbers based on specific brands and issuers.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Validate](#validate)
  - [Generate](#generate)
  - [Information](#information)
  - [Issue](#issue)
- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Troubleshooting](#troubleshooting)
- [Contributors](#contributors)
- [License](#license)

---

## Introduction

Credit cards play a vital role in modern commerce, with unique numbers identifying the cardholder and the issuing bank. This tool provides a way to:
- Validate credit card numbers using Luhn's Algorithm.
- Generate possible card numbers by replacing placeholders.
- Retrieve card details such as brand and issuer.
- Issue new valid card numbers for specific brands and issuers.

Credit card numbers follow specific rules:
- **Visa**: 13- and 16-digit numbers starting with 4.
- **MasterCard**: 16-digit numbers starting with 51-55.
- **American Express**: 15-digit numbers starting with 34 or 37.

The tool is designed to handle these rules and ensure compliance with proper formats and algorithms.

---

## Features

1. **Validate**  
   Check if a credit card number is valid using Luhn's Algorithm.

2. **Generate**  
   Generate possible credit card numbers by replacing placeholders (`*`) with digits.

3. **Information**  
   Retrieve details such as validity, brand, and issuer for a given card number.

4. **Issue**  
   Generate random valid card numbers for a specified brand and issuer.

---

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd creditcard
   ```

2. Build the project:
   ```bash
   go build -o creditcard .
   ```

---

## Usage

The tool offers four main features: `validate`, `generate`, `information`, and `issue`. Below are detailed instructions for each:

### Validate

Validates credit card numbers using Luhn's Algorithm.

#### Usage:
```bash
./creditcard validate <card_number> [<card_number> ...]
./creditcard validate --stdin
```

#### Examples:
```bash
$ ./creditcard validate "4400430180300003"
OK

$ ./creditcard validate "4400430180300002"
INCORRECT

$ echo "4400430180300003" | ./creditcard validate --stdin
OK
```

### Generate

Generates valid card numbers by replacing up to 4 asterisks (`*`) with digits. Supports random generation with the `--pick` flag.

#### Usage:
```bash
./creditcard generate <card_pattern>
./creditcard generate --pick <card_pattern>
```

#### Examples:
```bash
$ ./creditcard generate "440043018030****"
4400430180300003
4400430180300011
4400430180300029
...

$ ./creditcard generate --pick "440043018030****"
4400430180304385
```

In case of errors:
```bash
$ ./creditcard generate "44004301803*****"
$ echo $?
1
```

### Information

Fetches details about the card's validity, brand, and issuer based on `brands.txt` and `issuers.txt`.

#### Usage:
```bash
./creditcard information --brands=<brands_file> --issuers=<issuers_file> <card_number> [<card_number> ...]
./creditcard information --brands=<brands_file> --issuers=<issuers_file> --stdin
```

#### Examples:
```bash
$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
4400430180300003
Correct: yes
Card Brand: VISA
Card Issuer: Kaspi Gold

$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400450180300003"
4400450180300003
Correct: no
Card Brand: -
Card Issuer: -
```

### Issue

Generates a random valid card number for a specific brand and issuer.

#### Usage:
```bash
./creditcard issue --brands=<brands_file> --issuers=<issuers_file> --brand=<brand_name> --issuer=<issuer_name>
```

#### Example:
```bash
$ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
4400430180300003
```

---

## Configuration

The tool requires two files for the `information` and `issue` features:
1. `brands.txt`: Defines the mapping of card brands to their prefixes.
   Example:
   ```
   VISA:4
   MASTERCARD:51
   AMEX:34
   ```
2. `issuers.txt`: Defines the mapping of card issuers to their prefixes.
   Example:
   ```
   Kaspi Gold:440043
   Forte Black:404243
   ```

---

## Dependencies

This project is built using Go and adheres strictly to the following constraints:
- Only **built-in Go packages** are used.
- Code is formatted using `gofumpt`.

---

## Troubleshooting

1. **Compilation Issues**:  
   Ensure you are running the following command from the root directory:
   ```bash
   go build -o creditcard .
   ```

2. **Validation Errors**:  
   Check the input card number for length and format.

3. **Configuration Errors**:  
   Ensure the `brands.txt` and `issuers.txt` files are correctly formatted and accessible.

4. **Unhandled Panics**:  
   If unexpected errors occur, review your input format and file paths.

---

## Contributors

- **Koblandy Seipolla**  
    _Koweno_
  