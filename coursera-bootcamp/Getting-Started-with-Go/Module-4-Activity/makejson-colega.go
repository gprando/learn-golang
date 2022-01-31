package main
import ("fmt"
 "encoding/json"
 "bufio"
 "os"
 s "strings"
)

func main() {
  var input string
  result := make(map[string]string)
  fmt.Println("Please give a name: ")
  var reader bufio.Reader = *bufio.NewReader(os.Stdin)
  input, _ = reader.ReadString('\n')
  input = s.TrimSuffix(s.ToLower(input), "\n")

  result["name"] = input
  fmt.Println("Please give the address: ")
  var reader2 bufio.Reader = *bufio.NewReader(os.Stdin)
  input, _ = reader2.ReadString('\n')
  input = s.TrimSuffix(s.ToLower(input), "\n")
  result["address"] = input
  json, err := json.Marshal(result)
  if err != nil {
      fmt.Println("An error occured, " + err.Error())
  }
  fmt.Println(string(json))
}
