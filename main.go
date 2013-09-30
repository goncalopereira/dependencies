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


func CountsFromFileName(filename string, dependencies map[string]int) (depCount int) {
     if !strings.Contains(filename, ".dll") {
      return 0
    }

    val, ok := dependencies[filename]
 
    if ok == false {
      depCount = 1
    } else {
      depCount =  val+1
    }
  
    return
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

    depCount := CountsFromFileName(info.Name(), dependencies) 
    if depCount != 0 {
      dependencies[info.Name()] = depCount
    }
   
    return err 
   }

  err := filepath.Walk(tempRepository, wk)
  if err != nil {
    log.Fatal(err)
  }

  return dependencies
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
