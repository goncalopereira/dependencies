package git

import (
	"os"
	"os/exec"
)

func Clean(tempRepository string) {
  os.RemoveAll(tempRepository)
}

func Clone(repository, tempRepository string) error {
	cmd :=  exec.Command("git","clone",repository,tempRepository)
    
    err := cmd.Run()
   
    if err != nil {
   		return err
	}

	return nil   	
}
