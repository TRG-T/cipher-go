package main

import "fmt" 

var latin = [26]string{
	"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l",
	"m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x",
	"y", "z",
}

var galactic = [26]string{
	"ᔑ",  "ʖ",  "ᓵ",  "↸", "Ŀ", "⎓",
	"ㅓ", "〒", "╎",  "፧", "ꖌ", "ꖎ",
	"ᒲ",  "リ", "フ", "¡", "ᑑ", "።",
	"ነ",  "ﬧ",  "⚍", "⍊", "∴", "/",
	"॥",  "Λ",
}

func indexOf(text string) int {
	for i:=0; i<len(latin); i++ {
		if latin[i] == text {
			return i;
		}
	}
	return 0;
}

func reverse(text []string) []string {
	for i, a:= 0, len(text)-1; i<a; i, a = i+1, a-1 {
		text[i], text[a] = text[a], text[i];
	}
	return text;
}

func translateLetter(latinIndex int) string {
	return galactic[latinIndex]
}

func encrypt(text []string, key int) []string {
	for i:=0; i<len(text); i++ {
		latinIndex := indexOf(text[i]);
        if latinIndex+key > 26 {
            text[i] = latin[(latinIndex+key)%26];
        } else {
            text[i] = latin[latinIndex+key];
        }
		text[i] = translateLetter(indexOf(text[i]))
    }
    return text;
}

func main() {
	t := []string{"a", "b", "c", "z"};
	fmt.Printf("Original text: %v \n", t);
	reversedText := reverse(t);
	fmt.Printf("Reversed text: %v \n", reversedText);
	fmt.Printf("Encrypted with Caesar and translated into galactic: %v \n", encrypt(reversedText, 3));
}
