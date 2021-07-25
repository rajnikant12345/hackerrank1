/*
    https://www.hackerrank.com/contests/booking-hacakathon/challenges/similar-destinations/problem
*/

package main

import (
    "fmt"
    "sort"
    "strings"
)

var cityTagMap = map[string][]string{}
var uniqueCityTagMap = map[string]bool{}
var cityList []string
var brotherMap = map[string]bool{}

func findCommonTags(city1, city2 string) []string {
    var out []string
    for _, v := range cityTagMap[city1] {
        if uniqueCityTagMap[city2+"."+v] {
            out = append(out, v)
        }
    }
    return out
}

func checkTags(city string, tags []string) bool {
    count := 0
    for _, v := range tags {
        if uniqueCityTagMap[city+"."+v] {
            count++
        }
    }
    return count == len(tags)
}

func tagsAlreadyCatered(tags []string) bool {
    return brotherMap[strings.Join(tags, ",")]
}

type Result struct {
    Cities []string
    Tags   []string
}

func (r Result) Pack() string {
    str1 := strings.Join(r.Cities,",")
    str2 := strings.Join(r.Tags,",")
    return str1+":"+str2
}

func main() {
    common := 0
    fmt.Scanf("%d\n", &common)
    for {
        str := ""
        n, err := fmt.Scanf("%s", &str)
        if err != nil || n <= 0 {
            break
        }
        cityTags := strings.Split(str, ":")
        if len(cityTags) < 2 {
            continue
        }
        tags := strings.Split(cityTags[1], ",")
        cityList = append(cityList, cityTags[0])
        for _, v := range tags {
            cityTagMap[cityTags[0]] = append(cityTagMap[cityTags[0]], v)
            uniqueCityTagMap[cityTags[0]+"."+v] = true
        }
    }

    res := []Result{}

    for i := 0; i < len(cityList)-1; i++ {
        for j := i+1; j < len(cityList); j++ {
            var cl []string
            tags := findCommonTags(cityList[i], cityList[j])
            sort.Strings(tags)
            if tagsAlreadyCatered(tags) {
                continue
            }
            if len(tags) >= common {
                cl = append(cl, cityList[i], cityList[j])
                for k := 0; k < len(cityList); k++ {
                    if k == i || k == j {
                        continue
                    }
                    if checkTags(cityList[k], tags) {
                        cl = append(cl, cityList[k])
                    }
                }
            }
            if len(cl) > 0 {
                sort.Strings(cl)
                brotherMap[strings.Join(tags, ",")] = true
                res = append(res, Result{cl, tags})
            }
        }
    }

    sort.Slice(res, func(i, j int) bool {
        if len(res[i].Tags) == len(res[j].Tags) {
            str1 := res[i].Pack()
            str2 := res[j].Pack()
            res := strings.Compare(str1,str2)
            if res <= 0 {
                return true
            }
            return false
        }
        return len(res[i].Tags) > len(res[j].Tags)
    })
    str := ""
    for _,v := range res {
        str += fmt.Sprintln(v.Pack())
    }
    fmt.Println(str)
}
