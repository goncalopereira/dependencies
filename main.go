package main

import (
  "os"
  "fmt"
  "encoding/csv"
  "os/exec"
  "log"
  "path/filepath"
)

func Clean(tempRepository string) {
  os.RemoveAll(tempRepository)
}

func Clone(repository, tempRepository string) {
   cmd :=  exec.Command("git","clone",repository,tempRepository)
    
   err := cmd.Run()
   if err != nil {
    log.Fatal(err)
    }
  }


var wk = func(path string, info os.FileInfo, err error) error {
  fmt.Println(path)
  return err
}

func Files(tempRepository string) {  
  
  err := filepath.Walk(tempRepository, wk)
  if err != nil {
    log.Fatal(err)
  }
}


func main() {
 
  file, err := os.Open("projects.csv")
  
  if err != nil {
    return
  }
  
  defer file.Close()

  reader := csv.NewReader(file)

  lines, err := reader.ReadAll()
    
  repositories := make(map[string]string)

  for _, value := range lines {
    repositories[value[0]] = value[1]
  }
  

//  dependencies := make(map[string]map[string]int)

  tempRepository := os.TempDir() + "/tempRepo"
  for name, repository := range repositories {
    Clean(tempRepository)
    Clone(repository, tempRepository)
    fmt.Println(name)     
    Files(tempRepository)
  }

}
