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
  "bufio"
)

func Display(name string, dependencies map[string]int) {
    fmt.Println(name)
    for dll, count := range dependencies {
      fmt.Printf("%s %d\n", dll, count)   
  }
}

type Repository struct {
  dlls map[string]int
  usings map[string]int
}

func (r *Repository) DllCount(filename string) {
     if !strings.Contains(filename, ".dll") {
      return
    }

    filename = strings.TrimSuffix(filename, ".dll")

    val, ok := r.dlls[filename]
 
    if ok == false {
      r.dlls[filename] = 1
    } else {
      r.dlls[filename] = val+1
    }
  
    return
}

func (r *Repository) FileScan(path string) {
  if !strings.Contains(path,".cs") {
    return
  }

  file, err := os.Open(path)

  if err != nil {
    return
  }

  defer file.Close()

  scanner :=bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if strings.HasPrefix(line, "using ") {

      line = strings.TrimPrefix(line, "using ")
      line = strings.TrimSuffix(line,";")
    
      val, ok := r.usings[line]

      if ok == false {
        r.usings[line] = 1
      } else {
        r.usings[line] = val+1
      }
    }
  }
}

func Files(tempRepository string) (r Repository)  {  
  
  r = Repository {
    dlls: make(map[string]int),
    usings: make(map[string]int)}

  wk := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err 
    }

    if info.IsDir() {
      return nil
    }

    r.DllCount(info.Name())
    r.FileScan(path)
   
    return err 
  }

  err := filepath.Walk(tempRepository, wk)
  if err != nil {
    log.Fatal(err)
  }

  return 
}

func main() {
 
  repositoriesUrls, err := reader.ReadRepositories()
  
  if err != nil {
    log.Fatal(err)
  }

  for name, repository := range repositoriesUrls {
    tempRepository := os.TempDir() + "/tempRepo"
    git.Clean(tempRepository)
    err := git.Clone(repository, tempRepository)
    if err != nil {
      log.Fatal(err)
    }

    r := Files(tempRepository)       
    Display(name+"_dll", r.dlls)
    Display(name+"_usings", r.usings)
    writer.Write(name+"_dlls", r.dlls)
    writer.Write(name+"_usings", r.usings)

    if err != nil {
      log.Fatal(err)
    }
  }
}
