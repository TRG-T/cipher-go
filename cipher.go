package main

import (
	"fmt"
	"strings"
)

var latin = [5][6]string{
	{"a", "b", "c", "d", "e", "f"},
	{"g", "h", "i", "j", "k", "l"},
	{"m", "n", "o", "p", "q", "r"},
	{"s", "t", "u", "v", "w", "x"},
	{"y", "z", "!", "?", ":", " "},
}

var galactic = [5][6]string{	
	{"ᔑ", "ʖ", "ᓵ", "↸", "Ŀ", "⎓"},
	{"ㅓ", "〒", "╎", "፧", "ꖌ", "ꖎ"},
	{"ᒲ", "リ", "フ", "¡", "ᑑ", "።"},
	{"ነ", "ﬧ", "⚍", "⍊", "∴", "/"},
	{"॥", "Λ", "ʗ", "˨", "ᚴ", "ᚌ"},
}
var userText string
var userKey string

func indexOf(letter string) (int, int) {
	for i:=0; i<5; i++ {
		for a:=0; a<6; a++ {
			if latin[i][a] == letter {
				return i, a;
			}
		}
	}
	return 0, 0;
}

func reverse(text []string) []string {
	for i, a:= 0, len(text)-1; i<a; i, a = i+1, a-1 {
		text[i], text[a] = text[a], text[i];
	}
	return text;
}

func stringToKey(key []string) int {
	var keyInt int
	for i:=0; i<len(key); i++ {
		textRow, textColumn := indexOf(key[i])
		keyInt += textRow + textColumn
	}
	fmt.Println(keyInt)
	return keyInt
}

func encrypt(text []string, key int) string {
	for i:=0; i<len(text); i++ {
		latinRow, latinColumn := indexOf(text[i]);
		text[i] = galactic[latinRow][(latinColumn+key)%6]
    }
    return strings.Join(text, "");
}

func main() {
	fmt.Println("Enter text to encrypt:");
	fmt.Scanln(&userText)
	fmt.Println("Enter key:")
	fmt.Scanln(&userKey)
	key := stringToKey(strings.Split(userKey, ""));
	text := strings.Split(userText, "");
	fmt.Printf("Original text: %v \n", text);
	reversedText := reverse(text);
	fmt.Printf("Reversed text: %v \n", reversedText);
	fmt.Printf("Encrypted text: %v \n", encrypt(reversedText, key));
}
