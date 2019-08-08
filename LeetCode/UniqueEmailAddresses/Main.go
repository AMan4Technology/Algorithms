package UniqueEmailAddresses

import "strings"

func numUniqueEmails(emails []string) int {
	length := len(emails)
	if length < 2 {
		return length
	}
	var (
		countWithEmail = make(map[string]struct{})
		valid          strings.Builder
	)
	for _, email := range emails {
		for length, i := len(email), 0; i < length; i++ {
			if email[i] == '.' {
				continue
			}
			if email[i] == '+' {
				break
			}
			valid.WriteByte(email[i])
		}
		valid.WriteString(email[strings.IndexByte(email, '@'):])
		countWithEmail[valid.String()] = struct{}{}
		valid.Reset()
	}
	return len(countWithEmail)
}
