package main

import (
  "os"
  "fmt"
  "log"
  "path/filepath"
  "strings"
  "dependencies/git"
  "dependencies/writer"
  "dependencies/reader"
)

func Display(name string, dependencies map[string]int) {
    fmt.Println(name)
    for dll, count := range dependencies {
      fmt.Printf("%s %d\n", dll, count)   
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
 
  repositoriesUrls, err := reader.Read()
  
  if err != nil {
    log.Fatal(err)
  }

  dependencies := make(map[string]int)

  tempRepository := os.TempDir() + "/tempRepo"
  for name, repository := range repositoriesUrls {
    git.Clean(tempRepository)
    err := git.Clone(repository, tempRepository)

    if err != nil {
      log.Fatal(err)
    }

    dependencies = Files(tempRepository)       
    Display(name, dependencies)

    writer.Write(name, dependencies)

    if err != nil {
      log.Fatal(err)
    }
  }
}
