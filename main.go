package main

import (
  "fmt"
  "github.com/ttacon/chalk" //coloring console output
  "net/http" //http requests
  "io/ioutil"
  "strings" //for string manipulation
  "github.com/robertkrimen/otto" //javascript engine
)
func getLogic(file string) string {
  start := strings.Index(file, "logic{-") // finds "logic{* in the file"
  end := strings.Index(file, "-}") // finds "*}" in the file
  if start == -1 {
    js := "" // if no logic is provided sets the javascript to print "error no 'logic' element"
    return js
  } else{
    js := file[start + 7: end] // selects the contence of the logic element
    return js
  }
}
func getFileFromUrl() string { //function used to retrive a resource from url
  fmt.Print("Input URL: ")
  var url string
  fmt.Scanln(&url)
  file, _ := http.Get("http://" + url)
  defer file.Body.Close()
  contenceBytes, _ := ioutil.ReadAll(file.Body)
  contence := string(contenceBytes[:]) //convert a list of bytes to a string
  return contence
}
func splitUp(text string) []string{
  text1 := strings.Replace(text, "\"", "*", -1)
  a := strings.Split(text1, "*")
  return a
}
func display(file string){
  endOfJs := strings.Index(file, "-}")
  list := splitUp(file[endOfJs + 2:])
  js := getLogic(file)
  vm := otto.New() // creates the javascript VM
  vm.Run(js); //runs the javascript code in the vm
  fmt.Println("\t\t\t\t\t\t\t")
  fmt.Println("\t\t\t\t\t\t\t")
  end := len(list)
  i := 0
  for i < end {
    cur := list[i]
    if strings.HasPrefix(cur, "@red") {
      fmt.Print(chalk.Red, "")
    }
    if strings.HasPrefix(cur, "@blue") {
      fmt.Print(chalk.Blue, "")
    }
    if strings.HasPrefix(cur, "@green") {
      fmt.Print(chalk.Green, "")
    }
    if strings.HasPrefix(cur, "@yellow") {
      fmt.Print(chalk.Yellow, "")
    }
    if strings.HasPrefix(cur, "@white") {
      fmt.Print(chalk.White, "")
    }
    if strings.HasPrefix(cur, "@default") {
      fmt.Print("", chalk.Reset)
    }
    if strings.HasPrefix(cur, "@") == false && strings.HasPrefix(cur, "$") == false {
      fmt.Print(cur + " ")
    }
    if strings.HasPrefix(cur, "$") {
      value, _ := vm.Get(cur[1:])
      fmt.Print(value)
    }
    i++
  }
  // fmt.Print("\n\t\t\t\t\t\t\t");
  // fmt.Print("\n\t\t\t\t\t\t\t");
  // fmt.Print("\n\t\t\t\t\t\t\t");
  fmt.Println("\n\t\t\t\t\t\t\t", chalk.Reset);
}


func main() {
  file := getFileFromUrl()
  fmt.Println("\n\n\n")
  display(file)
  fmt.Print("> ")
  var input string
  fmt.Scanln(&input)
}
