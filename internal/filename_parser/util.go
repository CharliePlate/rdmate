package filename_parser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

type regexpList []*regexp2.Regexp

func compileRegexpList(rexpStrings []string) (regexpList, error) {
	rexpList := make(regexpList, 0, len(rexpStrings))

	for _, r := range rexpStrings {
		re, err := regexp2.Compile(r, regexp2.RE2)
		if err != nil {
			return nil, err
		}
		rexpList = append(rexpList, re)
	}

	return rexpList, nil
}

func mustCompileRegexpList(rexpStrings []string) regexpList {
	rexpList, err := compileRegexpList(rexpStrings)
	if err != nil {
		panic(err)
	}
	return rexpList
}

func (rl *regexpList) firstMatchingWithIndex(s string) (*regexp2.Regexp, int) {
	for i, r := range *rl {
		match, err := r.MatchString(s)
		if err != nil {
			continue
		}
		if match {
			return r, i
		}
	}
	return nil, -1
}

func (rl *regexpList) firstMatching(s string) *regexp2.Regexp {
	regexp, _ := rl.firstMatchingWithIndex(s)
	return regexp
}

func (rl *regexpList) replaceAllString(s string, repl string) string {
	for _, r := range *rl {
		rep, err := r.Replace(s, repl, 0, -1)
		if err != nil {
			continue
		}
		s = strings.TrimSpace(rep)
	}
	return s
}

func (rl *regexpList) parseFromPatterns(s string, options []string) string {
	_, index := rl.firstMatchingWithIndex(s)
	if index == -1 {
		return ""
	}
	return options[index]
}

func findAllGroups(re *regexp2.Regexp, s string) map[string]string {
	match, _ := re.FindStringMatch(s)
	if match == nil {
		return nil
	}

	matchMap := make(map[string]string)
	for _, group := range match.Groups() {
		if group.Length > 0 && group.Name != "" {
			matchMap[group.Name] = group.String()
		}
	}
	return matchMap
}
