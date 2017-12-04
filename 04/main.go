package main

// To ensure security, a valid passphrase must contain no duplicate words.
//
// --- Part Two ---
// For added security, yet another system policy has been put in place. Now, a
// valid passphrase must contain no two words that are anagrams of each other -
// that is, a passphrase is invalid if any word's letters can be rearranged to
// form any other word in the passphrase.
//
// Answer part one: 383
// Answer part two: 265

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	fmt.Println("Valid passphrases:", ValidPassphrases(file, false))

	_, err = file.Seek(0, 0)
	fmt.Println("Valid passphrases (part two):", ValidPassphrases(file, true))
}

func ValidPassphrases(f *os.File, anagram bool) int {
	var valid int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())

		dup := map[string]bool{}
		for _, v := range r {
			if anagram {
				// sort string - easier to compare
				v = sortString(string(v))
			}

			dup[v] = true
		}

		if len(dup) == len(r) {
			valid++
		}
	}

	return valid
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// eof
