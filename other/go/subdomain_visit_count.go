package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func main() {
	// 	testList := []string{"9001 leetcode.com", "9001 discuss.leetcode.com"}
	// 	fmt.Println(subdomainVisits(testList))
	doms := map[string]int{}
	doms["test"] = 1
	doms["test"] += 3
	fmt.Println(doms["test"])
}

// func subdomainVisits(cpdomains []string) []string {
// 	var results []string
// 	domainVisits := map[string]int{}
// 	for _, cp := range cpdomains {
// 		count, domains := splitCountPair(cp)
// 		for _, d := range domains {
// 			if _, ok := domainVisits[d]; !ok {
// 				c, _ := strconv.Atoi(count)
// 				domainVisits[d] += c
// 			}
// 		}
// 	}

// 	for k, v := range domainVisits {
// 		strVal := strconv.Itoa(v)
// 		result := strings.Join([]string{strVal, k}, " ")
// 		results = append(results, result)
// 	}
// 	return results
// }

// func splitCountPair(cp string) (string, []string) {
// 	cpSplit := strings.Split(cp, " ")
// 	count := cpSplit[0]
// 	domain := cpSplit[1]
// 	return count, splitSubDomains(domain)
// }

// func splitSubDomains(domain string) []string {
// 	var result []string
// 	subs := strings.Split(domain, ".")
// 	for i := 0; i < len(subs); i++ {
// 		result = append(result, strings.Join(subs[i:], "."))
// 	}
// 	return result
// }
