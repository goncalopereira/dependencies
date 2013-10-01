package reader

import (
	"encoding/csv"
	"os"
  "path/filepath"
  "strings"
)

func FillDependencies(path string, d map[string]bool) (dependencies map[string]int, allDependencies map[string]bool) {
  allDependencies = d
  dependencies = make(map[string]int)

  return
  
}

func ReadCSV(extension string) (dependencies map[string]map[string]int, allDependencies map[string]bool, err error) {
  dependencies = make(map[string]map[string]int)
  allDependencies = make(map[string]bool)

  wk := func(path string, info os.FileInfo, err error) error {

    project := strings.TrimSuffix(info.Name(),extension)

    dependencies[project], allDependencies = FillDependencies(path, allDependencies)       

    return err
  }

  err = filepath.Walk("outputs",wk)

  return  
}
 
func ReadRepositories() (map[string]string, error) {

  file, err := os.Open("projects.csv")
  
  if err != nil {
    return nil, err
  }
  
  defer file.Close()

  reader := csv.NewReader(file)

  lines, err := reader.ReadAll()
    
  repositoriesUrls := make(map[string]string)

  for _, value := range lines {
    repositoriesUrls[value[0]] = value[1]
  }


  return repositoriesUrls, nil
}
