package ach

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Validator struct{}

func (v *Validator) parseNumField(r string) (s int) {
	s, err := strconv.Atoi(strings.TrimSpace(r))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	return s
}

func (v *Validator) formatFileCreationDate(t time.Time) string {
	return t.Format("060102")
}

func (v *Validator) parseFileCreationDate(s string) time.Time {
	t, _ := time.Parse("060102", s)
	return t
}

func (v *Validator) formatFileCreationTime(t time.Time) string {
	return t.Format("1504")
}

func (v *Validator) parseFileCreationTime(s string) time.Time {
	t, _ := time.Parse("1504", s)
	return t
}

func (v *Validator) isUpperAlphanumeric(s string) (b bool) {
	if regexp.MustCompile(`[^A-Z0-9]+`).MatchString(s) {
		return false
	}
	return true
}

func (v *Validator) isAlphanumeric(s string) (b bool) {
	if regexp.MustCompile(`[^ a-zA-Z0-9_*\/]+`).MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return false
	}
	return true
}

func (v *Validator) IsLetter(s string) bool {
	fmt.Println(s)
	for _, r := range s {
		fmt.Printf("%v ", r)
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func (v *Validator) rightPad(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + (overallLen - len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func (v *Validator) leftPad(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + (overallLen - len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
