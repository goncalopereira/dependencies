package reader

import (
	"encoding/csv"
	"os"
  "path/filepath"
  "strings"
  "strconv"
  "log"
)

func FillDependencies(path string, d map[string]bool) (dependencies map[string]int, allDependencies map[string]bool) {
  allDependencies = d
  dependencies = make(map[string]int)
  
  file, err := os.Open(path)
  if err != nil {
    return
  }
  
  defer file.Close()

  reader := csv.NewReader(file)

  lines, err := reader.ReadAll()

  for _, value := range lines {

    if strings.Contains(value[0],",") {
      log.Fatal("dependency " + value[0] + " contains comma ','")
    }
    i, err := strconv.Atoi(value[1])
  
    if err == nil {
      dependencies[value[0]] = i
    }
    
    if !allDependencies[value[0]] {    
      allDependencies[value[0]] = true
    }
  }
   
  return
  
}

func ReadCSV(extension string) (dependencies map[string]map[string]int, allDependencies map[string]bool, err error) {
  dependencies = make(map[string]map[string]int)
  allDependencies = make(map[string]bool)

  wk := func(path string, info os.FileInfo, err error) error {
        
    if strings.Contains(info.Name(), extension) && !strings.Contains(info.Name(), "result") {
    
      log.Println(info.Name())    

      project := strings.TrimSuffix(info.Name(),extension)
    
      dependencies[project], allDependencies = FillDependencies(path, allDependencies)       
    }

    return err
  }

  err = filepath.Walk("output",wk)

  return  
}
 
func ReadRepositories(filename string) (map[string]string, error) {

  file, err := os.Open(filename)
  
  if err != nil {
    return nil, err
  }
  
  defer file.Close()

  reader := csv.NewReader(file)

  lines, err := reader.ReadAll()
    
  repositoriesUrls := make(map[string]string)

  for _, value := range lines {
    name := strings.Replace(value[0],"/","-",-1)
    name = strings.Replace(name,".","-",-1)
    repositoriesUrls[name] = value[1]
  }


  return repositoriesUrls, nil
}
