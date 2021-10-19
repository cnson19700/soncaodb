package namestand

import (
	"log"
	"regexp"
	"strings"
)

func Check(str string) string {
	re1, err := regexp.Compile(`[^a-zA-Z\s]+`)

	if err != nil {
		log.Fatal(err)
	}

	re2, err := regexp.Compile(`\s\s+`)

	if err != nil {
		log.Fatal(err)
	}

	re3, err := regexp.Compile(`^ | $`)

	if err != nil {
		log.Fatal(err)
	}

	match := re1.ReplaceAllString(str, "")     // remove all digit and non-word
	match2 := re2.ReplaceAllString(match, " ") // remove double space
	match3 := re3.ReplaceAllString(match2, "") // remove space in first and end
	match3 = strings.ToLower(match3)
	strArr := strings.Split(match3, "")
	strArr[0] = strings.ToUpper(strArr[0])

	for i := 0; i < len(strArr); i++ {
		if strArr[i] == " " {
			strArr[i+1] = strings.ToUpper(strArr[i+1])
		}
	}

	return strings.Replace(strings.Join(strArr[:], ","), ",", "", -1)
}
