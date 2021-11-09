package main

import "fmt" 

var latin = [26]byte{
	'a', 'b', 'c', 'd', 'e', 'f',
	'g', 'h', 'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p', 'q', 'r',
	's', 't', 'u', 'v', 'w', 'x',
	'y', 'z',
}

var galactic = [26]rune{
	'ᔑ',  'ʖ',  'ᓵ',  '↸', 'Ŀ', '⎓',
	'ㅓ', '〒', '╎',  '፧', 'ꖌ', 'ꖎ',
	'ᒲ',  'リ', 'フ', '¡', 'ᑑ', '።',
	'ነ',  'ﬧ',  '⚍', '⍊', '∴', '/',
	'॥',  'Λ',
}

func reverse(text []byte) []byte {
	for i, a:= 0, len(text)-1; i<a; i,a = i+1, a-1 {
		text[i], text[a] = text[a], text[i]
	}
	return text
}

func cesarEncrypt(text []byte) []byte {
	for i:=0; i<len(text); i++ {
        if(text[i]+3 > 122) {
            text[i] -= 22;
        } else {
            text[i] += 3;
        }
    }
    return text;
}

func main() {
	t := []byte{'a', 'b', 'c'}
	fmt.Printf("Oryginalny tekst: %v \n", string(t))
	reversedText := reverse(t)
	fmt.Printf("Odwrócony tekst: %v \n", string(reversedText))
	fmt.Printf("Zaszyfrowany cezarem: %v \n", string(cesarEncrypt(reversedText)))
}
