package main

import "fmt"

func OgDefaultImage() string {
	return fmt.Sprintf("%s/favicon-32x32.png", BaseURL())
}
