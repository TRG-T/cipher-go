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


func main() {
	t := []byte{'a', 'b', 'c'}
	fmt.Printf("Oryginalny tekst: %v \n", string(t))
	fmt.Printf("Odwrócony tekst: %v \n", string(reverse(t)))
}
