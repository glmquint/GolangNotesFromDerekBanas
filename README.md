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

### runes

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

var pl = fmt.Println

func main() {
    rStr := "abcdefg"
    pl("Rune Count :", utf8.RuneCountInString(rStr))
    for i, runeVal := range rStr {
        fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
    }
}
```

### time

```go
    now := time.Now()
    pl(now.Year(), now.Month(), now.Day())
    pl(now.Hour(), now.Minute(), now.Second())
```

### random values

```go
package main

import (
    "fmt"
    "time"
    "math/rand"
)

var pl = fmt.Println

func main() {
    seedSecs := time.Now().Unix()
    rand.Seed(seedSecs)
    randNum := rand.Intn(50) + 1
    pl("Random :", randNum)
}
```

### math

Most common math operators (after `import("math")`)

- `math.Abs`
- `math.Pow`
- `math.Sqrt`
- `math.Cbrt`
- `math.Ceil`
- `math.Floor`
- `math.Round`
- `math.Log2`
- `math.Log10`
- `math.Max`
- `math.Min`

`math.Pi`

### for loops

```go
    // for initialization; condition; postStatement {BODY}
    for x := 1; x <= 5; x++ {
        pl(x)
    }
```

### while loops

```go
for true {
    if done {
        break
    } else if skip {
        continue
    }
    // do stuff
}
```

### range

```go
    aNums :=  []int{1, 2, 3}
    for _, num := range aNums {
        pl(num)
    }
```

### arrays

> A collection of values with the same data-type

You are allowed to change the values in an array, but you're not allowed to change its size

```go
    // basic declaration
    var arr1 [5]int
    arr1[0] = 1

    // declaration and initialization
    arr2 := [5]int{1, 2, 3, 4, 5}

    // index accessing
    pl("Index 0 :", arr2[0])

    // array length
    pl("Arr Length :", len(arr2))

    for i:= 0; i<len(arr2); i++ {
        pl(arr2[i])
    }
    for i, v := range arr2 {
        fmt.Printf("%d : %d\n", i, v)
    }

    // multidimensional array
    arr3 := [2][2]int{
        {1, 2},
        {3, 4},
    }
    for i:=0; i<2; i++{
        for j:=0; j<2; j++{
            pl(arr3[i][j])
        }
    }
```

```go
    aStr1 := "abcde"
    rArr := []rune(aStr1)
    for _, v := range rArr {
        fmt.Printf("Rune Array : %d\n", v)
    }
    byteArr := []byte{'a', 'b', 'c'}
    bStr := string(byteArr[:])
    pl("I'm a string :", bStr)
```

### slices

They're like arrays but the **can** grow

```go
    // declaration
    // var name []datatype
    sl1 := make([]string, 6)

    // initialization
    sl1[0] = "Society"
    sl1[1] = "of"
    sl1[2] = "the"
    sl1[3] = "Simulated"
    sl1[4] = "Universe"
    pl("Slice Size :", len(sl1))
    for i:=0; i<len(sl1); i++ {
        pl(sl1[i])
    }
    for _, x := range sl1 {
        pl(x)
    }

    // slice of an array
    sArr := [5]int{1, 2, 3, 4, 5}
    sl3 := sArr[0:2]
    pl("1st 3 values :", sArr[:3])
    pl("from 3rd :", sArr[2:])
    pl("last val :", sArr[len(sArr)-1]) // sArr[-1]
    pl("last 3 :", sArr[len(sArr)-2:]) // sArr[-2:]

    //changing the slice will also change the array
    sArr[0] = 10
    pl("sl3 :", sl3)
    sArr[0] = 1
    pl("sl3 :", sl3)

    // appending
    sl3 = append(sl3, 12)
    pl("sl3 :", sl3)
    pl("sArr :", sArr)

    // empty string as default value for slices
    sl4 := make([]string, 6)
    pl("sl4 :", sl4)
    pl("sl4[0] :", sl4[0])
```

### functions

```go
func sayHello() {
    pl("Hello")
}

func main() {
    // func funcName(parameters) returnType {BODY}
    sayHello()

}
```

```go
func getSum(x int, y int) int {
    return x + y
}

func main() {
    pl(getSum(5, 4))
}
```

```go
func getTwo(x int) (int, int) {
    return x+1, x+2
}

func main() {
    pl(getTwo(5))
}
```

```go
func getQuotient(x float64, y float64) (ans float64, err error) {
    if y == 0 {
        return 0, fmt.Errorf("You can't divide by zero")
    } else {
        return x / y, nil
    }
}

func main() {
    pl(getQuotient(5, 4))
    pl(getQuotient(5, 0))
}
```

```go
func getSum2(nums ...int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func main() {
    pl(getSum2(1, 2, 3, 4, 5))
}
```

```go
func getArraySum(arr []int) int {
    sum := 0
    for _, val := range arr {
        sum += val
    }
    return sum
}

func main() {
    vArr := []int{1, 2, 3, 4}
    // pass-by-value
    pl("Array Sum :", getArraySum(vArr))
}
```

```go

func changeVal(f3 int) int {
    f3 += 1
    return f3
}

func main() {
    f3 := 5
    pl("f3 before func :", f3) // 5
    changeVal(f3)
    pl("f3 after func :", f3) // 5
}
```

```go

func changeVal(myPtr *int) {
    *myPtr = 12
}

func main() {
    f3 := 5
    pl("f3 before func :", f3) // 5
    changeVal(&f3) // pass-by-reference
    pl("f3 after func :", f3) // 12
}
```

```go
func changeVal(myPtr *int) {
    *myPtr = 12
}

func main() {
    f4 := 10
    var f4Ptr *int = &f4
    pl("f4 Address :", f4Ptr)
    pl("f4 Value :", *f4Ptr) // 10
    *f4Ptr = 11
    pl("f4 Value :", *f4Ptr) // 11

    pl("f4 before func :", f4) // 11
    changeVal(&f4) // pass-by-reference
    pl("f4 after func :", f4) // 12
}
```

```go
func dblArrVals(arr *[4]int) {
    for x := 0; x < 4; x++ {
        arr[x] *= 2
    }
}

func main() {
    pArr := [4]int{1, 2, 3, 4}
    dblArrVals(&pArr)
    pl(pArr) // 2 4 6 8
}
```

```go
func getAverage(nums ...float64) float64 {
    var sum float64 = 0.0
    var numSize float64 = float64(len(nums))
    for _, val := range nums {
        sum += val
    }
    return (sum / numSize)
}

func main() {
    iSlice := []float64{11, 13, 17}
    fmt.Printf("Avarage : %.3f\n", getAverage(iSlice...)) // 13.667
}
```

### file IO

File creation:
```go
    f, err := os.Create("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
```

`defer` will be executed when `f` goes out-of-scope

File write:
```go
    for _, num := range sPrimeArr {
        _, err := f.WriteString(num + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
```

File read:
```go
    scan1 := bufio.NewScanner(f)
    for scan1.Scan() {
        pl("Prime :", scan1.Text())
    }
    if err := scan1.Err(); err != nil {
        log.Fatal(err)
    }
```

File append:
```go
    /*
    Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified

    O_RDONLY    : open the file read-only
    O_WRONLY    : open the file write-only
    O_RDWR      : open the file read-write

    These can be or'ed

    O_APPEND    : append data to the file when writing
    O_CREATE    : create a new fileif none exists
    O_EXCL      : used with O_CREATE, file must not exists
    O_SYNC      : open for synchronous I/O
    O_TRUNC     : truncate regural writable file when opened
    */

    _, err := os.Stat("data.txt")
    if errors.Is(err, os.ErrNotExist) {
        pl("File Dosn't Exists")
    } else {
        f, err := os.OpenFile("data.txt",
            os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
        if _, err := f.WriteString("13\n"); err != nil {
            log.Fatal(err)
        }
    }
```

### Command Line
