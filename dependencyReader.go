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

func (r *Repository) FileScan(path string) (wholeFile string) {
  if !strings.HasSuffix(path,".cs") && !strings.HasSuffix(path,".vb") && !strings.HasSuffix(path,".config") && !strings.Contains(path,".asp") {
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

    wholeFile += path+","+line+"\n"

    if strings.HasPrefix(line, "Imports ") {
      fmt.Println(line) 
     line = strings.TrimPrefix(line, "Imports ")
    
      val, ok := r.usings[line]

      if ok == false {
        r.usings[line] = 1
      } else {
        r.usings[line] = val+1
      }
    

    } else if strings.HasPrefix(line, "using ") {

      line = strings.TrimPrefix(line, "using ")
      line = strings.TrimSuffix(line,";")
      line_split := strings.Split(line," = ")
      line = line_split[len(line_split)-1] 
    
      val, ok := r.usings[line]

      if ok == false {
        r.usings[line] = 1
      } else {
        r.usings[line] = val+1
      }
    }
  }

  return
}

func Files(name, tempRepository string) (r Repository)  {  

  file, err := os.Create("output/code_" + name+ ".csv")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
 
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
    wholeFile := r.FileScan(path)
    file.WriteString(wholeFile)
   
    return err 
  }

  err = filepath.Walk(tempRepository, wk)
  if err != nil {
    log.Fatal(err)
  }

  return 
}

func Execute(name, repository, dlls, usings string) error {
    
    log.Println(name)   
    tempRepository := "tempRepo"  
    log.Println("clean")
    git.Clean(tempRepository)
    log.Println("clone")
    out, err := git.Clone(repository, tempRepository)
    if err != nil {
      log.Printf("%s\n",out)
      return err
    }

    log.Println(tempRepository)
    r := Files(name, tempRepository)       
    
    writer.Write(dlls, r.dlls)
    writer.Write(usings, r.usings)

    return err
}

func main() {
 
  repositoriesUrls, err := reader.ReadRepositories("projects.csv")
  
  if err != nil {
    log.Fatal(err)
  }

  for name, repository := range repositoriesUrls {

   dlls := "./output/" + name +"_dlls.csv"
   usings := "./output/" + name + "_usings.csv"
   
    _, err := os.Stat(dlls)
   
    log.Println(err) 
    
    if err != nil {
        err = Execute(name, repository, dlls, usings)
        if err != nil {
          log.Fatal(err)
        }
    }
  }
}
