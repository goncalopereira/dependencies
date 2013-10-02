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

func SortProjects(dependencies map[string]map[string]int) []string {
  sorted := make([]string, len(dependencies))

  i := 0
  for proj, _ := range dependencies {
    sorted[i] = proj
    i++
  }

  sort.Strings(sorted)
  return sorted
}

func matrix(extension string) {

  dependencies, allDependencies, err := reader.ReadCSV(extension)

  sortedKeys := SortDependencies(allDependencies)
  sortedProjects := SortProjects(dependencies)
  
  if err != nil {
    log.Fatal(err)
  }
  
  matrix := make([][]string, len(allDependencies)+1)
  for i := 0; i < len(allDependencies)+1;i++ {
    matrix[i] = make([]string, len(dependencies)+1)
  }

  matrix[0][0] = "x"
 
  projId  := 1
  for _, projName := range sortedProjects {
    matrix[0][projId] = projName
    depId :=1
    for _, dependency := range sortedKeys {
      matrix[depId][0] = dependency
      matrix[depId][projId] = strconv.Itoa(dependencies[projName][dependency])
      depId++
    }
    projId++
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

