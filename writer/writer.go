package writer

import (
	"os"
	"encoding/csv"
	"strconv"
  "log"
)

func Write(name string, dependencies map[string]int) error {
    
    file, err := os.Create(name)
    if err != nil {
      return err
    }
  
    log.Println(name)

    defer file.Close()
 
    writer := csv.NewWriter(file)

    aDependencies := make([][]string, len(dependencies))

    i:=0
    for key, value := range dependencies {
      row := make([]string, 2)
      row[0] = key
      row[1] = strconv.Itoa(value)
      aDependencies[i] = row

      i++ 
    }

    err = writer.WriteAll(aDependencies)

    return err 
}
