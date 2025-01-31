package services

import (
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/labstack/echo"
)

// Anagram ...
type AnagramChecker struct {
	WordOne   string `json:"wordOne"`
	WordTwo   string `json:"wordTwo"`
	IsAnagram bool   `json:"isAnagram"`
	Message   string `json:"message"`
}

// CheckAnagram ...
func CheckAnagram(c echo.Context) error {

	ac := &AnagramChecker{}

	ac.WordOne = strings.ToLower(c.FormValue("wordOne"))
	ac.WordTwo = strings.ToLower(c.FormValue("wordTwo"))
	ac.IsAnagram = false

	// Input validations
	var isStringAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString

	if !isStringAlphanumeric(ac.WordOne) || !isStringAlphanumeric(ac.WordTwo) {
		ac.Message = "Both words have to be alphanumeric"
		data := echo.Map{
			"isAnagram": ac.IsAnagram,
			"message":   ac.Message,
		}
		return c.Render(http.StatusBadRequest, "checkAnagram.html", data)
	}

	if len(ac.WordOne) < 3 || len(ac.WordOne) > 24 {
		ac.Message = "Length of word one must be greater than 3 and less than 24 characters"
		data := echo.Map{
			"isAnagram": ac.IsAnagram,
			"message":   ac.Message,
		}
		return c.Render(http.StatusBadRequest, "checkAnagram.html", data)
	}

	if len(ac.WordTwo) < 3 || len(ac.WordTwo) > 24 {
		ac.Message = "Length of word two must be greater than 3 and less than 24 characters"
		data := echo.Map{
			"isAnagram": ac.IsAnagram,
			"message":   ac.Message,
		}
		return c.Render(http.StatusBadRequest, "checkAnagram.html", data)
	}

	// Anagram check
	wordOneSlice := strings.Split(ac.WordOne, "")
	wordTwoSlice := strings.Split(ac.WordTwo, "")

	sort.Slice(wordOneSlice, func(i, j int) bool {
		return wordOneSlice[i] < wordOneSlice[j]
	})

	sort.Slice(wordTwoSlice, func(i, j int) bool {
		return wordTwoSlice[i] < wordTwoSlice[j]
	})

	var sortedWordOne string = strings.Join(wordOneSlice, "")
	var sortedWordTwo string = strings.Join(wordTwoSlice, "")

	if sortedWordOne == sortedWordTwo {
		ac.IsAnagram = true
		ac.Message = "Yep, those words are anagrams of each other"
	} else {
		ac.IsAnagram = false
		ac.Message = "Nope, these words are not anagrams"
	}

	data := echo.Map{
		"message":   ac.Message,
		"isAnagram": ac.IsAnagram,
	}
	return c.Render(http.StatusOK, "checkAnagram.html", data)
}
