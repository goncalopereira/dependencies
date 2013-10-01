package git

import (
	"os"
	"os/exec"
)

func Clean(tempRepository string) {
  os.RemoveAll(tempRepository)
}

func Clone(repository, tempRepository string) ([]byte, error) {
	cmd :=  exec.Command("git","clone","--depth","1",repository,tempRepository)
    
    out, err := cmd.Output()
   
    if err != nil {      
   		return out, err
	}

	return nil, nil
}

