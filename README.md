# GOLANG tutorial


Following [this](https://www.youtube.com/watch?v=YzLrWHZa-Kc) tutorial by Derek Banas

## Installation of a new version

Follow instructions in [here](https://go.dev/doc/install)

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
|-----------|---------------|
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


Accessible with `os.Args` as an array of strings

## Packages


In a new directory (`app`) we use the command `go mod init example/project` to create a new go module named `example/project`

> Go modules allow us to manage libraries. 
> They contain one project (or library) and a collection of go packages

We create a new `main.go` file, making sure that it's at the same level of the generated `app/go.mod` file

We create a package as a file in a new subdirectory in our module: `app/mypackage/mypackage.go`

(Make sure every name is in lowercase)
Every object in Uppercase will be available outside of out package

In the `main.go` app we use 
```go
import(
    stuff "example/project/mypackage"
)
```
to import the package `mypackage` (and all its exported variables and methods) from the `stuff` module

We can simply use them as `fmt.Println("Hello", stuff.Name)` or `strArr := stuff.IntArrToStrArr(intArr)`


## Maps


> Collection of key-value pairs.
> Key can be any data-type that can be compared using ==

```go
    // var myMap map [keyType]valueType
    var heroes map[string]string
    heroes = make(map[string]string)

    villains := make(map[string]string)
    heroes["Batman"] = "Bruce Wayne"
    heroes["Superman"] = "Clark Kent"
    heroes["The Flash"] = "Barry Allen"

    villains["Lex Luther"] = "Lex Luther"

    superPets := map[int]string{1: "Krypto",
        2:"Bat Houd"}
    fmt.Printf("Batman is %v\n", heroes["Batman"])
    pl("Chip : ", superPets[3])
    _, ok := superPets[3]
    pl("Is there a 3rd pet :", ok)
    for k, v := range heroes {
        fmt.Printf("%s is %s\n", k, v)
    }
    delete(heroes, "The Flash")
```

## Generics


> With Generics we can specify the datatype to be used at a later time
> It's mainly used when we want functions that can operate on multiple datatypes

(Also, newly added in 1.18)

```go
type MyConstraint interface {
    int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
    return x+y
}

func main() {
    pl("5 + 4 =", getSumGen(5, 4))
    pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))
}
```

## Struct


> They allow to store values with many different datatypes in a very structured way

```go
type customer struct {
    name string
    address string
    bal float64
}

func getCustomerInfo(c customer) {
    fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
    c.address = address
}

func main() {
    var tS customer
    tS.name = "Tom Smith"
    tS.address = "5 main st"
    tS.bal = 234.56
    getCustomerInfo(tS)
    newCustAdd(&tS, "123 South st")
    pl("Address :", tS.address)
    sS := customer{"Sally Smith", "123 Main", 0.0}
    pl("Name :", sS.name)
}
```

A function **only** for the struct `rectangle`:
```go
func (r rectangle) Area() float64 {
    return r.length * r.height
}
```

> Go doesn't support inheritance but _does_ support composition by embedding a struct inside another struct

```go
type contact struct {
    fName string
    lName string
    phone string
}

type business struct {
    name string
    address string
    contact
}

func (b business) info(){
    fmt.Printf("Contact at %s is %s %s", b.name,
        b.contact.fName, b.contact.lName)
}

func main() {
    con1 := contact {
        "James",
        "Wang",
        "555-1212",
    }
    bus1 := business{
        "ABC Plumbering",
        "234 North St",
        con1,
    }
    bus1.info()

}
```

## Defined Types


```go
type Tsp float64
type Tbs float64
type ML float64

func tspToML(tsp Tsp) ML {
    return ML(tsp * 4.92)
}

func tbsToML(tsp Tsp) ML {
    return ML(tsp * 14.79)
}

func main() {
    ml1 := ML(Tsp(3) * 4.92)
    fmt.Printf("4 tsps = %.2f ML\n", ml1)
    ml2 := ML(Tbs(3) * 14.79)
    fmt.Printf("3 TBs = %.2f ML\n", ml2)

    pl("2 tsp + 4 tsp = ", Tsp(2), Tsp(4))
    pl("2 tsp > 4 tsp = ", Tsp(2) > Tsp(4))

    fmt.Printf("3 tsp = %.2f Ml\n", tspToML(3))
    fmt.Printf("3 tbs = %.2f Ml\n", tbsToML(3))

}
```

But there's a better way to do it!

### Associate Methods


```go
type Tsp float64
type Tbs float64
type ML float64

func (tsp Tsp) toMLs() ML {
    return ML(tsp * 4.92)
}

func (tbs Tbs) toMLs() ML {
    return ML(tbs * 14.79)
}

func main() {
    tsp1 := Tsp(3)
    fmt.Printf("%.2f tsp = %.2f ML\n", tsp1, tsp1.toMLs())
}
```

## Protecting Data 


### Encapsulation


In our package we create a struct and some getter and setter functions to access the private data (lowercase components of struct)
```go
type Date struct{
    day int
    month int
    year int
}

// Setters
func (d *Date) SetDay(day int) error {
    if (day < 1) || (day > 31) {
        return errors.New("incorrect day value")
    }
    d.day = day
    return nil
}
func (d *Date) SetMonth(m int) error {
    if (m < 1) || (m > 12) {
        return errors.New("incorrect month value")
    }
    d.month = m
    return nil
}
func (d *Date) SetYear(y int) error {
    if (y < 1875) || (y > time.Now().Year()) {
        return errors.New("incorrect year value")
    }
    d.year = y
    return nil
}

// Getters
func (d *Date) Day() int {
    return d.day
}
func (d *Date) Month() int {
    return d.month
}
func (d *Date) Year() int {
    return d.year
}
```

Then, in the main app, we use them as follows:
```go
    date := stuff.Date{}
    err := date.SetMonth(12)
    if err != nil {
        log.Fatal(err)
    }
    err = date.SetDay(21)
    if err != nil {
        log.Fatal(err)
    }
    err = date.SetYear(1974)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("1st Day : %d/%d/%d\n",
        date.Month(), date.Day(), date.Year())
```

### Interfaces


> Interfaces allow us to create contracts that say that if anything inherit them, they must implement some predefined methods

You don't actually have to define that a type is gonna use a specifi interface. All you need to to is implement the functions and everything is automatic (that the type is gonna implement the interface)

## Concurrency (GoRoutines)


Simply add `go` before a function and it will operate on a different thread (goroutine)

```go
func printTo15(){
    for i := 1; i <= 15; i++ {
        pl("Fun 1 :", i)
    }
}

func printTo10(){
    for i := 1; i <= 10; i++ {
        pl("Fun 2 :", i)
    }
}

func main() {
    go printTo15()
    go printTo10()
    time.Sleep(2*time.Second)
}
```

You need to put a sleep on the main thread, or all its children will be terminated before they can complete

Being on different threads, you can no longer expect in what order this functions will execute

### Channels


> Channels allow different go routines to comunicate

Channels can only carry values of **one** type

```go
func nums1(channel chan int){
    channel <- 1
    channel <- 2
    channel <- 3
}

func nums2(channel chan int){
    channel <- 4
    channel <- 5
    channel <- 6
}

func main() {
    channel1 := make(chan int)
    channel2 := make(chan int)
    go nums1(channel1)
    go nums2(channel2)
    pl(<-channel1)
    pl(<-channel1)
    pl(<-channel1)
    pl(<-channel2)
    pl(<-channel2)
    pl(<-channel2)
}

```
Now the values will always be in order

### Mutex


A lock for mutually exclusive access operations

```go
type Account struct{
    balance int
    lock sync.Mutex
}
func (a *Account) GetBalance() int {
    a.lock.Lock()
    defer a.lock.Unlock()
    return a.balance
}
func (a *Account) Withdraw(v int) {
    a.lock.Lock()
    defer a.lock.Unlock()
    if v > a.balance {
        pl("Not enough money in account")
    } else {
        fmt.Printf("%d withdrawn : Balance : %d\n",
            v, a.balance)
        a.balance -= v
    }
}

func main() {
    var acct Account
    acct.balance = 100
    pl("Balance :", acct.GetBalance())
    for i:=0; i<12; i++ {
        go acct.Withdraw(10)
    }
    time.Sleep(2*time.Second)
}
```

## Clojures


> Clojures are functions that don't have to be associated with an identifier, but often are associated with variables

```go
    intSum := func(x, y int) int { return x + y }
    pl("5 + 4 =", intSum(5, 4))

    samp1 := 1
    changeVar := func() {
        samp1 += 1
    }
    changeVar()
    pl("samp1 =", samp1)
```

Note how clojures **can** change values outside of the function

### Passing functions


You can pass functions to a function

```go
func useFunc(f func(int, int) int, x, y int){
    pl("Answer :", (f(x, y)))
}

func sumvals(x, y int) int {
    return x + y
}

func main() {
    useFunc(sumvals, 5, 8) // 13
}
```

## Regex


```go
package main

import (
    "fmt"
    "regexp"
)

var pl = fmt.Println

func main() {
    // regex matching
    reStr := "The ape was at the apex"
    match, _ := regexp.MatchString("ape[^ ]", reStr)
    pl(match)

    // compiled expression
    reStr2 := "Cat rat mat fat pat"
    r, _ := regexp.Compile("[crmfp]at")
    pl("MatchString :", r.MatchString(reStr2))
    pl("FindString :", r.FindString(reStr2))
    pl("Index :", r.FindStringIndex(reStr2))
    pl("All String :", r.FindAllString(reStr2, -1))
    pl("First 2 Strings", r.FindAllString(reStr2, 2))
    pl("All Submatch Index :", r.FindAllStringSubmatchIndex(reStr2, -1))
    pl(r.ReplaceAllString(reStr2, "Dog"))
}
```

## Automated Testing


> Tests make sure that the code continues to work as we're developing it

In a new go module (app2, after `go mod init app2`) we have two files:
- `testemail.go` with the functionality that we're developing
- `testemail_test.go` with the testcases

```go:testemail.go
package app2

import (
    "fmt"
    "regexp"
)

func IsEmail(s string) (string, error) {
    r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}\.[A-Za-z]{2,3}`)
    if r.MatchString(s) {
        return "Valid Email", nil
    } else {
        return "", fmt.Errorf("not a valid email")
    }
}

```

```go:testemail_test.go
package app2
import("testing")

func TestIsEmail(t *testing.T){
    _, err := IsEmail("hello")
    if err == nil {
        t.Error("hello is not an email") // PASS
    }
    _, err = IsEmail("glmquint@gmail.com")
    if err != nil {
        t.Error("glmquint@gmail.com is an email") // FAIL
    }
    _, err = IsEmail("glmquint@gmail")
    if err == nil {
        t.Error("glmquint@gmail is not an email") // FAIL
    }
}
```

In the app2 directory we can then use the command `go test -v` to execute all tests in the `*_test.go` file

