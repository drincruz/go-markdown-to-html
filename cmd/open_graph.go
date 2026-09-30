package main

import "fmt"

func OgDefaultImage() string {
	return fmt.Sprintf("%s/favicon-32x32.png", BaseURL())
}

func OgDefaultDescription() string {
	return "Software engineer, sometimes speaker, always a learner."
}
