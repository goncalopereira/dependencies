package main

import (
  "os"
  "fmt"
  "encoding/csv"
  "os/exec"
  "log"
  "path/filepath"
  "strings"
)

func Display(dependencies map[string]map[string]int) {
  for name, dep := range dependencies {
    fmt.Println(name)
    for dll, count := range dep {
      fmt.Printf("%s %d\n", dll, count)
    }
  }
}

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

func Files(tempRepository string) (dependencies map[string]int) {  
  
  dependencies = make(map[string]int)  

  wk := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err 
    }

    if info.IsDir() {
      return nil
    }
    
    if !strings.Contains(info.Name(), ".dll") {
      return nil
    }

     val, ok := dependencies[info.Name()]
     
    if ok == false {
      dependencies[info.Name()] = 1
    } else {
      dependencies[info.Name()] = val+1
    }
   
    return err 
   }

  err := filepath.Walk(tempRepository, wk)
  if err != nil {
    log.Fatal(err)
  }

  return
}

func main() {
 
  file, err := os.Open("projects.csv")
  
  if err != nil {
    return
  }
  
  defer file.Close()

  reader := csv.NewReader(file)

  lines, err := reader.ReadAll()
    
  repositoriesUrls := make(map[string]string)

  for _, value := range lines {
    repositoriesUrls[value[0]] = value[1]
  }

  
  dependencies := make(map[string]map[string]int)

  tempRepository := os.TempDir() + "/tempRepo"
  for name, repository := range repositoriesUrls {
    Clean(tempRepository)
    Clone(repository, tempRepository)
    dependencies[name] = Files(tempRepository)       
  }

  Display(dependencies)

}
