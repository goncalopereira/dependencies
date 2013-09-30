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

type Repository struct {
  dlls map[string]int
}

func (r *Repository) DllCount(filename string) {
     if !strings.Contains(filename, ".dll") {
      return
    }

    val, ok := r.dlls[filename]
 
    if ok == false {
      r.dlls[filename] = 1
    } else {
      r.dlls[filename] = val+1
    }
  
    return
}

func Files(tempRepository string) (r Repository)  {  
  
  r = Repository {
    dlls: make(map[string]int)}

  wk := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err 
    }

    if info.IsDir() {
      return nil
    }

    r.DllCount(info.Name())
   
    return err 
  }

  err := filepath.Walk(tempRepository, wk)
  if err != nil {
    log.Fatal(err)
  }

  return 
}

func main() {
 
  repositoriesUrls, err := reader.Read()
  
  if err != nil {
    log.Fatal(err)
  }

  tempRepository := os.TempDir() + "/tempRepo"
  for name, repository := range repositoriesUrls {
    git.Clean(tempRepository)
    err := git.Clone(repository, tempRepository)

    if err != nil {
      log.Fatal(err)
    }

    r := Files(tempRepository)       
    Display(name, r.dlls)

    writer.Write(name, r.dlls)

    if err != nil {
      log.Fatal(err)
    }
  }
}
