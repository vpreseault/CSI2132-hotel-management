package main

import (
	"log"
	"os"

	"github.com/caarlos0/env/v10"
)

type ccors struct {
	AllowedOrigins []string `env:"ALLOWED_ORIGINS,required,notEmpty" envSeparator:","`
}

var (
	Cors = &ccors{}
)

func init() {
	toParse := []any{Cors}
	errors := []error{}

	for _, v := range toParse {
		if err := env.Parse(v); err != nil {
			if er, ok := err.(env.AggregateError); ok {
				errors = append(errors, er.Errors...)

				continue
			}

			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		log.Fatal("errors found while parsing environment variables")

		os.Exit(1)
	}
}
