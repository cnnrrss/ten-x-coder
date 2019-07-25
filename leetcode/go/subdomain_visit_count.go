func subdomainVisits(cpdomains []string) []string {
    var results []string
    domainVisits := map[string]int{}
    for _, cp := range cpdomains {
        count, domains := splitCountPair(cp)
        for _, d := range domains {
            c, _ := strconv.Atoi(count)
            domainVisits[d] += c
        }
    }

    for k, v := range domainVisits {
        strVal := strconv.Itoa(v)
        result := strings.Join([]string{strVal, k}, " ")
        results = append(results, result)
    }
    return results
}

func splitCountPair(cp string) (string, []string) {
    cpSplit := strings.Split(cp, " ")
    count := cpSplit[0]
    domain := cpSplit[1]
    return count, splitSubDomains(domain)
}

func splitSubDomains(domain string) []string {
    var result []string
    subs := strings.Split(domain, ".")
    for i := 0; i < len(subs); i++ {
        result = append(result, strings.Join(subs[i:], "."))
    }
    return result
}