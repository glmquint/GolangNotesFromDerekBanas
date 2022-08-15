package main

import (
    "fmt"
    "regexp"
)

var pl = fmt.Println

func main() {
    reStr := "The ape was at the apex"
    match, _ := regexp.MatchString("ape[^ ]", reStr)
    pl(match)

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
