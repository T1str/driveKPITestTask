package validators

import (
	"regexp"
	"testtask/internal/entities"
)

func ValidateRequest(req entities.Request) bool {
	return validateField(req.PeriodStart) &&
		validateField(req.PeriodEnd) &&
		validateField(req.PeriodKey) &&
		validateField(req.IndicatorToMoID) &&
		validateField(req.IndicatorToMoFactID) &&
		validateField(req.Value) &&
		validateField(req.FactTime) &&
		validateField(req.IsPlan) &&
		validateField(req.AuthUserID) &&
		validateField(req.Comment)
}

func validateField(field string) bool {
	if isEmpty(field) {
		return false
	}
	return !detectSecurityVulnerabilities(field)
}

func detectSecurityVulnerabilities(input string) bool {
	return detectSQLInjection(input) ||
		detectPathTraversal(input) ||
		detectXSS(input) ||
		detectSSRF(input)
}

func isEmpty(input string) bool {
	return input == ""
}

func detectSQLInjection(input string) bool {
	pattern := `[;'"\\/*%<>()]`
	return regexp.MustCompile(pattern).MatchString(input)
}

func detectPathTraversal(input string) bool {
	pattern := `\.\./|\\\.\.\\`
	return regexp.MustCompile(pattern).MatchString(input)
}

func detectXSS(input string) bool {
	pattern := `</|<|>|/`
	return regexp.MustCompile(pattern).MatchString(input)
}

func detectSSRF(input string) bool {
	pattern := `http://|https://`
	return regexp.MustCompile(pattern).MatchString(input)
}
