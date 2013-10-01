package main

import (
  "dependencies/reader"  
  "log"
  "sort"
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

func main() {
  dependencies, allDependencies, err := reader.ReadCSV("_dlls.csv")

  sortedKeys := SortDependencies(allDependencies)

  if err != nil {
    log.Fatal(err)
  }
  
  for key, deps := range dependencies {
    log.Println(key)
    
    for dep, value := range deps {
      log.Printf("%s %d\n", dep, value)
    }
  }

  log.Println("all")
  for _,val := range sortedKeys {
    log.Println(val)
  }
  
}

