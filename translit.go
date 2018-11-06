package main

var trTable = map[rune]rune{
	// cyrillic
	'q':  'я',
	'w':  'ш',
	'e':  'е',
	'r':  'р',
	't':  'т',
	'y':  'ы',
	'u':  'у',
	'i':  'и',
	'o':  'о',
	'p':  'п',
	'[':  'ю',
	']':  'ж',
	'a':  'а',
	's':  'с',
	'd':  'д',
	'f':  'ф',
	'g':  'г',
	'h':  'ч',
	'j':  'й',
	'k':  'к',
	'l':  'л',
	'\\': 'э',
	'`':  'щ',
	'z':  'з',
	'x':  'х',
	'c':  'ц',
	'v':  'в',
	'b':  'б',
	'n':  'н',
	'm':  'м',
	'-':  'ь',
	'=':  'ъ',
	'Q':  'Я',
	'W':  'Ш',
	'E':  'Е',
	'R':  'Р',
	'T':  'Т',
	'Y':  'Ы',
	'U':  'У',
	'I':  'И',
	'O':  'О',
	'P':  'П',
	'{':  'Ю',
	'}':  'Ж',
	'A':  'А',
	'S':  'С',
	'D':  'Д',
	'F':  'Ф',
	'G':  'Г',
	'H':  'Ч',
	'J':  'Й',
	'K':  'К',
	'L':  'Л',
	'|':  'Э',
	'~':  'Щ',
	'Z':  'З',
	'X':  'Х',
	'C':  'Ц',
	'V':  'В',
	'B':  'Б',
	'N':  'Н',
	'M':  'М',
	'_':  'Ь',
	'+':  'Ъ',
	// latin
	'я': 'q',
	'ш': 'w',
	'е': 'e',
	'р': 'r',
	'т': 't',
	'ы': 'y',
	'у': 'u',
	'и': 'i',
	'о': 'o',
	'п': 'p',
	'ю': '[',
	'ж': ']',
	'а': 'a',
	'с': 's',
	'д': 'd',
	'ф': 'f',
	'г': 'g',
	'ч': 'h',
	'й': 'j',
	'к': 'k',
	'л': 'l',
	'э': '\\',
	'щ': '`',
	'з': 'z',
	'х': 'x',
	'ц': 'c',
	'в': 'v',
	'б': 'b',
	'н': 'n',
	'м': 'm',
	'ь': '-',
	'ъ': '=',
	'Я': 'Q',
	'Ш': 'W',
	'Е': 'E',
	'Р': 'R',
	'Т': 'T',
	'Ы': 'Y',
	'У': 'U',
	'И': 'I',
	'О': 'O',
	'П': 'P',
	'Ю': '{',
	'Ж': '}',
	'А': 'A',
	'С': 'S',
	'Д': 'D',
	'Ф': 'F',
	'Г': 'G',
	'Ч': 'H',
	'Й': 'J',
	'К': 'K',
	'Л': 'L',
	'Э': '"',
	'Щ': '~',
	'З': 'Z',
	'Х': 'X',
	'Ц': 'C',
	'В': 'V',
	'Б': 'B',
	'Н': 'N',
	'М': 'M',
	'Ь': '_',
	'Ъ': '+',
}

var ruKl = "com.apple.keylayout.Russian-Phonetic"

// var enKl = "com.apple.keylayout.US"
var enKl = "com.apple.keylayout.ABC"

// var issw = "issw"

func tr(s string) string {
	str := ""
	for _, c := range s {
		if t, ok := trTable[c]; ok {
			str += string(t)
		} else {
			str += string(c)
		}
	}
	return str
}

/*
func main() {
	// localPart = os.Getenv("LOCAL_PART")
	issw := os.Getenv("HOME") + "/bin/issw"
	// str, err := ioutil.ReadAll(os.Stdin)
	query := ""

	if len(os.Args) > 1 {
		query = os.Args[1]
	} else {
		os.Exit(-1)
	}

	out, err := exec.Command(issw).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	cuKl := strings.TrimSpace(string(out))
	// fmt.Println(strings.TrimSpace(string(cuKl)))
	// fmt.Println(cuKl)
	if cuKl == enKl {
		_, err := exec.Command(issw, ruKl).Output()
		if err != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err)
		}
		//fmt.Println(string(out))
	} else {
		_, err := exec.Command(issw, enKl).Output()
		if err != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err)
		}
		//fmt.Println(string(out))
	}

	fmt.Println(tr(string(query)))
	os.Exit(0)

}
*/
