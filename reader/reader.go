package reader

import (
	"encoding/csv"
	"os"
)

func Read() (map[string]string, error) {

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