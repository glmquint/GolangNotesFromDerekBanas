# GOLANG tutorial

Following [this](https://www.youtube.com/watch?v=YzLrWHZa-Kc) tutorial by Derek Banas

---

## Syntax

Simplest Hello World:

```go
package main

import (
    "fmt"
)

var pl = fmt.Println

func main() {
    pl("Hello Go")
}
```

(Also, I've added this line to .vimrc: `autocmd FileType go map <F3> <Esc>:w<CR>:!clear && go run %<CR>`)

std-input example with error handling with log module:

```go
package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
)

var pl = fmt.Println

func main() {
    pl("What is your name?")
    reader := bufio.NewReader(os.Stdin)
    name, err := reader.ReadString('\n')
    if err == nil {
        pl("Hello", name)
    } else {
        log.Fatal(err)
    }
}
```

```go
package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
)

var pl = fmt.Println

func main() {
    // var name type
    var vName string = "Ghi0m"
    var v1, v2 = 1.2, 3.4
    var v3 = "hello"
    // variables are mutable by default
    v4 := 2.4
    // v4 := 5.4 // this would be a compilation error:
    // redefinition on var v4
    // Instead, use
    v4 = 5.4
}
```

Available variable types are:

| type      | default value |
|===========|===============|
| int       | 0             |
| float64   | 0.0           |
| bool      | false         |
| string    | ""            |
| rune      | ''            |

```go
package main

import (
    "fmt"
    "reflect"
)

var pl = fmt.Println

func main() {
    pl(reflect.TypeOf(25))      // int
    pl(reflect.TypeOf(3.14))    // float64
    pl(reflect.TypeOf(true))    // bool
    pl(reflect.TypeOf("Ghi0m")) // string
    pl(reflect.TypeOf('😎'))    // int32

}
```

### casting
```go
    cV1:=1.5
    cV2:= int(cV1)
    pl(cV2)
```

```go
    cV3 := "50000000"
    cV4, err := strconv.Atoi(cV3)
    pl(cV4, err, reflect.TypeOf(cV4)) // 50000000 <nil> int
```

```go
    cV5 := 500000
    cV6 := strconv.Itoa(cV5) // 500000
    pl(cV6)
```

`strconv.ParseFloat("3.14", 64)`

`fmt.Sprintf("%f", 3.14)`

### if conditional

```go
    iAge := 8
    if (iAge >= 1) && (iAge <= 18) {
        pl("Important Birthday")
    } else if (iAge == 32) || (iAge == 50) {
        pl("Important Birthday")
    } else if iAge >= 65 {
        pl("Important Birthday")
    } else {
        pl("Not an important Birthday")
    }
    pl("!true = ", !true)
```

### strings

```go
    sV1 := "A word" // an array of bytes
    replacer := strings.NewReplacer("A", "Another")
    sV2 := replacer.Replace(sV1)
    pl(sV2)
    pl("Length :", len(sV2))
    pl("Contains Another :", strings.Contains(sV2, "Another"))
    pl("o index :", strings.Index(sV2, "i"))
    pl("Replace :", strings.Replace(sV2, "o", "0", -1))
    sV3 := "\nSome Words\n" // \t \"
    sV3 = strings.TrimSpace(sV3)
    pl("Split :", strings.Split("a-b-c-d", "-"))
    pl("Lower :", strings.ToLower(sV2))
    pl("Upper :", strings.ToUpper(sV2))
    pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
    pl("Prefix :", strings.HasPrefix("tacocat", "cat"))
```

#checkpoint @28:23
