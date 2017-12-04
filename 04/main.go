package main

// To ensure security, a valid passphrase must contain no duplicate words.
//
// --- Part Two ---
// For added security, yet another system policy has been put in place. Now, a
// valid passphrase must contain no two words that are anagrams of each other -
// that is, a passphrase is invalid if any word's letters can be rearranged to
// form any other word in the passphrase.

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

	fmt.Println("Valid passphrases:", ValidPassphrases(file))
	_, err = file.Seek(0, 0)
	fmt.Println("Valid passphrases:", ValidPassphrasesAnagrams(file))
}

func ValidPassphrases(f *os.File) int {
	var valid int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())

		dup := map[string]bool{}
		for k, _ := range r {
			if dup[r[k]] != true {
				dup[r[k]] = true
			}
		}

		if len(dup) == len(r) {
			valid++
		}
	}

	return valid
}

func ValidPassphrasesAnagrams(f *os.File) int {
	var valid int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())
		fmt.Println(r)

		dup := map[string]int{}
		for _, v := range r {
			// sort string - easier to compare
			s := sortString(string(v))
			fmt.Println(s)

			dup[s]++
		}

		// if lenghts match there's no duplicates
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
