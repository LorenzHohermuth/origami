package cmd

import "fmt"

func binary(s string) string {
    res := ""
    for _, c := range s {
        res = fmt.Sprintf("%s%.8b", res, c)
    }
    return res
}
