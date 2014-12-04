package config

import (
	"regexp"
)

const IpAddressPattern = `([0-9]{1,3}\.){3}[0-9]{1,3}`
const UuidPattern = `([0-9a-z]{32})|([0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12})`

func IsUuid(value string) bool {
	matched, err := regexp.MatchString(UuidPattern, value)
	if err != nil {
		return false
	}
	return matched
}
