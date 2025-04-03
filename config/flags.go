package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Command string

	// For validate
	CardNumbersToValidate []string
	Stdin    	bool

	// For generate
	CardNumberToGenerate string
	Pick		bool
	
	// For information
	BrandsFile string
	IssuersFile string
	CardNumbersToInform []string
	
	// For issue
	Brand string
	Issuer string
}

var Cfg Config

func ParseCommand() error {
	args := os.Args[1:]
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments")
	}

	command := args[0]
	os.Args = append(os.Args[:1], os.Args[2:]...)
	switch command {
	case "validate":
		Cfg.Command = command
		flag.BoolVar(&Cfg.Stdin, "stdin", false, "To get card numbers from standard input pipeline")
		flag.Parse()
		args = flag.Args()
		if !Cfg.Stdin {
			Cfg.CardNumbersToValidate = args
		} else if len(args) > 2{
			return fmt.Errorf("extra arguments with --stdin")
		}
	case "generate":
		Cfg.Command = command 
		flag.BoolVar(&Cfg.Pick, "pick", false, "To generate only one random credit card")
		flag.Parse()
		args = flag.Args()
		fmt.Println("THIS IS ARGUMENTS: ", args)
		if !Cfg.Pick  &&  len(args) != 1{
			return fmt.Errorf("expected only one credit card number prompt")
		} else if len(args) != 1 && Cfg.Pick{
			return fmt.Errorf("expected only one card number after --pick")
		} else if !Cfg.Pick {
			
		} else if Cfg.Pick {

		}
	case "information":
		Cfg.Command = command
		flag.StringVar(&Cfg.BrandsFile, "brands", "./data/brands.txt", "Path to file with brands of cards")
		flag.StringVar(&Cfg.IssuersFile, "issuers", "./data/issuers.txt", "Path to file with issuers of cards")
		flag.Parse()
		args = flag.Args()
		if len(args) < 1 {
			return fmt.Errorf("at least one card number must be given")
		}
		Cfg.CardNumbersToInform = args
	case "issue":
		Cfg.Command = command
		flag.StringVar(&Cfg.BrandsFile, "brands", "./data/brands.txt", "Path to file with brands of cards")
		flag.StringVar(&Cfg.IssuersFile, "issuers", "./data/issuers.txt", "Path to file with issuers of cards")
		flag.StringVar(&Cfg.Brand, "brand", "", "Brand to issue card number")
		flag.StringVar(&Cfg.Issuer, "issuer", "", "Issuer to issue card number")
		flag.Parse()
		args = flag.Args()
		if len(args) > 0 {
			return fmt.Errorf("extra unknown flags for issue")
		}
	default: return fmt.Errorf("unknown command")
	}

	flag.Usage = func() {fmt.Println(HelpMessage)}

	return nil
}

var HelpMessage = `Usage: creditcard [command] [options] [arguments]

A program for credit card validation, generation, and information retrieval.

Commands:
  validate     Validate credit card number(s) using Luhn's Algorithm.
               - Supports multiple numbers as arguments.
               - Use --stdin to read numbers from standard input.

  generate     Generate possible credit card numbers by replacing up to 4 trailing asterisks (*) with digits.
               - Use --pick to randomly select a single valid number.
               - Exits with status 1 for invalid input.

  information  Retrieve details about credit card(s), including validity, brand, and issuer.
               - Requires --brands and --issuers files for lookup.
               - Supports multiple numbers as arguments.
               - Use --stdin to read numbers from standard input.

  issue        Generate a random valid credit card number for a specified brand and issuer.
               - Requires --brands and --issuers files.
               - Specify --brand and --issuer for targeted generation.
               - Exits with status 1 for errors.

Options:
  --help   Show this help message and exit.
  --stdin      Read credit card numbers from standard input (for validate and information commands).
  --pick       Randomly select a single result (for generate command).
  --brands     Path to the file containing card brand prefixes (for information and issue commands).
  --issuers    Path to the file containing issuer prefixes (for information and issue commands).
  --brand      Specify the brand for card issuance (for issue command).
  --issuer     Specify the issuer for card issuance (for issue command).`
