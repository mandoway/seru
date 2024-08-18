package language

import (
	"cuelang.org/go/cue/scanner"
	"cuelang.org/go/cue/token"
	"log"
)

func CountTokensUsingScanner(source []byte) int {
	eh := func(_ token.Pos, msg string, args []interface{}) {
		log.Println("Error during token count:", msg)
	}

	// verify scan
	var s scanner.Scanner
	s.Init(token.NewFile("", -1, len(source)), source, eh, scanner.ScanComments|scanner.DontInsertCommas)

	count := 0

	for {
		_, tok, _ := s.Scan()
		if tok == token.EOF {
			break
		}
		if s.ErrorCount > 0 {
			return -1
		}

		count++
	}

	return count
}
