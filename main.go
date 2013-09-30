package main

import (
  "os"
  "fmt"
  "encoding/csv"
)

func Clean(tempRepository string) {
  fmt.Println("delete " + tempRepository)  
}

func Clone(repository, tempRepository string) {
  command := "git clone " + repository + " " + tempRepository
  fmt.Println(command)
}

func main() {
 
  file, err := os.Open("projects.csv")
  
  if err != nil {
    return
  }
  
  defer file.Close()

  reader := csv.NewReader(file)

  lines, err := reader.ReadAll()
    
  repositories := make(map[string]string)

  for _, value := range lines {
    repositories[value[0]] = value[1]
  }
  

//  dependencies := make(map[string]map[string]int)

  tempRepository := os.TempDir() + "/tempRepo"
  for name, repository := range repositories {
    Clean(tempRepository)
    Clone(repository, tempRepository)
    fmt.Println(name)     
  }

}
