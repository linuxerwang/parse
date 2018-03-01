package parse

import (
	"io/ioutil"
	"log"
)

// InfoLogger is used to print informational message, default to off
var InfoLogger = log.New(ioutil.Discard, "", 0)
