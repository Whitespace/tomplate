package main

import (
  "io/ioutil"
  "log"
  "gopkg.in/yaml.v2"
  "os"
  "text/template"
)

func main() {
  // Usage info
  if len(os.Args) < 3 {
    log.Print("Usage: tomplate values.yaml template.txt")
    return
  }

  // Read in data
  valuesfile := os.Args[1]
  valuesdata, err := ioutil.ReadFile(valuesfile)
  if err != nil { log.Fatal(err) }

  // Injest data
  var values map[string]interface{}
  err = yaml.Unmarshal(valuesdata, &values)
  if err != nil { log.Fatal(err) }

  // Read in template
  templatefile := os.Args[2]
  templatedata, err := ioutil.ReadFile(templatefile)
  if err != nil { log.Fatal(err) }

  // Validate template
  tmpl, err := template.New("output").Parse(string(templatedata))
  if err != nil { log.Fatal(err) }

  // Apply data to template
  err = tmpl.Execute(os.Stdout, values)
  if err != nil { log.Fatal(err) }
}
