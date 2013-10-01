package main

import (
  "dependencies/reader"  
  "log"
  "sort"
  "os"
  "encoding/csv"
  "strconv"
)

func SortDependencies(dependencies map[string]bool) []string {
  sorted := make([]string, len(dependencies))
  
  i := 0
  for dep, _ := range dependencies {
    sorted[i] = dep
    i++
  }

  sort.Strings(sorted)
  return sorted
}

func matrix(extension string) {

  dependencies, allDependencies, err := reader.ReadCSV(extension)

  sortedKeys := SortDependencies(allDependencies)

  if err != nil {
    log.Fatal(err)
  }
  
  matrix := make([][]string, len(allDependencies)+1)
  
  //projects
  matrix[0] = make([]string, len(dependencies)+1)
  matrix[0][0] = "x"
  
  projId := 1
  for projName, _ := range dependencies {
        matrix[0][projId] = projName
        projId++
  }

  depId :=1
  
  for _, dep := range sortedKeys {
    matrix[depId] = make([]string, len(dependencies)+1)
    matrix[depId][0] = dep 
    
    projId = 1
    for projName, _ := range dependencies {
      val, ok := dependencies[projName][dep]
      if ok == true {
        matrix[depId][projId] = strconv.Itoa(val)
      } else {
        matrix[depId][projId] = "0"
      }
      projId++
    }
    
    depId++
  }

  file, err := os.Create("output/" + extension+ "_results.csv")
    if err != nil {
      log.Fatal(err)
  }

    defer file.Close()
 
    writer := csv.NewWriter(file)

    writer.WriteAll(matrix)
}

func main() {
  matrix("_dlls.csv")
  matrix("_usings.csv")
}

