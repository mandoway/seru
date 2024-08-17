package syntactic

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func ReduceSyntactically(reductionCmd *exec.Cmd) error {
	// Todo add time measurement

	log.Println("Running command:", reductionCmd)

	reductionCmd.Stderr = os.Stderr
	err := reductionCmd.Run()
	if err != nil {
		return errors.New("syntactic reduction failed: " + err.Error())
	}

	return nil
}
