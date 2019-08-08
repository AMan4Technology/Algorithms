package SubdomainVisitCount

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) (result []string) {
	counts := make(map[string]int, len(cpdomains))
	for _, value := range cpdomains {
		offset := strings.IndexByte(value, ' ')
		count, _ := strconv.Atoi(value[:offset])
		for domain := value[offset+1:]; domain != ""; domain = domain[offset+1:] {
			counts[domain] += count
			if offset = strings.IndexByte(domain, '.'); offset == -1 {
				break
			}
		}
	}
	result = make([]string, 0, len(counts))
	var sb strings.Builder
	for domain, count := range counts {
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(' ')
		sb.WriteString(domain)
		result = append(result, sb.String())
		sb.Reset()
	}
	return
}
