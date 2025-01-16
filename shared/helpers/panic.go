package helpers

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
