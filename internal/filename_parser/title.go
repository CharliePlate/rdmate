package filename_parser

import (
	"fmt"
	"regexp"
	"strings"
)

var regexpTitle = mustCompileRegexpList([]string{
	// Special, Despecialized, etc. Edition Movies, e.g: Mission.Impossible.3.Special.Edition.2011
	`(?i)^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*\(?\b(?<edition>(((Extended.|Ultimate.)?(Director.?s|Collector.?s|Theatrical|Anniversary|The.Uncut|Ultimate|Final(?=(.(Cut|Edition|Version)))|Extended|Rogue|Special|Despecialized|\d{2,3}(th)?.Anniversary)(.(Cut|Edition|Version))?(.(Extended|Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit))?|((Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit|Edition|Restored|((2|3|4)in1))))))\b\)?.{1,3}(?<year>(1(8|9)|20)\d{2}(?!p|i|\d+|\]|\W\d+)))+(\W+|_|$)(?!\\)`,
	// Folder movie format, e.g: Blade Runner 2049 (2017)
	`(?i)^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*\((?<year>(1(8|9)|20)\d{2}(?!p|i|(1(8|9)|20)\d{2}|\]|\W(1(8|9)|20)\d{2})))+`,
	// Normal movie format, e.g: Mission.Impossible.3.2011
	`(?i)^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*(?<year>(1(8|9)|20)\d{2}(?!p|i|(1(8|9)|20)\d{2}|\]|\W(1(8|9)|20)\d{2})))+(\W+|_|$)(?!\\)`,
	// PassThePopcorn Torrent names: Star.Wars[PassThePopcorn]
	`(?i)^(?<title>.+?)?(?:(?:[-_\W](?<![()[!]))*(?<year>(\[\w *\])))+(\W+|_|$)(?!\\)`,
	// That did not work? Maybe some tool uses [] for years. Who would do that?
	`(?i)^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)!]))*(?<year>(1(8|9)|20)\d{2}(?!p|i|\d+|\W\d+)))+(\W+|_|$)(?!\\)`,
	// As a last resort for movies that have ( or [ in their title.
	`(?i)^(?<title>.+?)?(?:(?:[-_\W](?<![)[!]))*(?<year>(1(8|9)|20)\d{2}(?!p|i|\d+|\]|\W\d+)))+(\W+|_|$)(?!\\)`,
})

type Title struct {
	Title   string
	Year    string
	Edition string
}

func ParseTitle(filename string) Title {
	first := regexpTitle.firstMatching(filename)
	if first == nil {
		return Title{}
	}

	groups := findAllGroups(first, filename)
	return Title{
		Title:   simplifyTitle(groups["title"]),
		Year:    groups["year"],
		Edition: groups["edition"],
	}
}

var regexpSimplifyTitle = mustCompileRegexpList([]string{
	`(?i)\s*(?:480[ip]|576[ip]|720[ip]|1080[ip]|2160[ip]|HVEC|[xh][\W_]?26[45]|DD\W?5\W1|[<>?*:|]|848x480|1280x720|1920x1080)((8|10)b(it))?`,
	`(?i)^\[\s*[a-z]+(\.[a-z]+)+\s*\][- ]*|^www\.[a-z]+\.(?:com|net)[ -]*`,
	`(?i)^\[(?:REQ)\]`,
	`(?i)\[(?:ettv|rartv|rarbg|cttv)\]$`,
	`(?i)\b(Bluray|(dvdr?|BD)rip|HDTV|HDRip|TS|R5|CAM|SCR|(WEB|DVD)?.?SCREENER|DiVX|xvid|web-?dl)\b`,
	`(?i)\[.+?\]`,
	`(?i)\b((Extended.|Ultimate.)?(Director.?s|Collector.?s|Theatrical|Anniversary|The.Uncut|DC|Ultimate|Final(?=(.(Cut|Edition|Version)))|Extended|Special|Despecialized|unrated|\d{2,3}(th)?.Anniversary)(.(Cut|Edition|Version))?(.(Extended|Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit))?|((Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit|Edition|Restored|((2|3|4)in1)))){1,3}`,
	`(?i)\b(TRUE.?FRENCH|videomann|SUBFRENCH|PLDUB|MULTI)`,
	`(?i)\b(PROPER|REAL|READ.NFO)`,
})

func simplifyTitle(title string) string {
	s := regexp.MustCompile("_").ReplaceAllString(title, " ")
	s = regexpSimplifyTitle.replaceAllString(s, "")

	languages := parseLanguages(s)
	for _, lang := range languages {
		s = strings.ReplaceAll(s, lang, "")
	}

	var result strings.Builder
	parts := strings.Split(s, ".")

	for n := 0; n < len(parts); n++ {
		part := parts[n]
		nextPart := safeGetIndex(parts, n+1)

		for len(part) == 1 && len(nextPart) == 1 {
			result.WriteString(fmt.Sprintf("%s.", part))
			n++
			part = safeGetIndex(parts, n)
			if len(part) > 1 {
				result.WriteString(" ")
			}
		}

		result.WriteString(part)
		result.WriteString(" ")
	}

	return strings.TrimSpace(result.String())
}

func safeGetIndex[T any](arr []T, i int) T {
	if i >= len(arr) {
		var zeroValue T
		return zeroValue
	}
	return arr[i]
}
