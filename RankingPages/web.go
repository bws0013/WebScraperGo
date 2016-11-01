package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
	"golang.org/x/net/html"
	"net/http"
	"sort"
	"strings"
)

/*
	This is going to add the ability to see how many times a particular page is linked to.
	Essentially instead of seeing the links found on pages the id-dfs will occur then a
	list of links along with their number of occurrences will be returned in sorted order.
*/

// The number of times a particular page has been linked
var linkNumbers map[string]int

// TODO: sort the values of linkNumbers before printing them.
// TODO: remove non http/s links before doing anything else.
func main() {

	linkNumbers = make(map[string]int)

	// url is the starting point of our search
	url := "http://auburn.edu/~bws0013/"

	//links := getPageWords(url)

	// provide our url and the maximum depth we are trying to go within the web pages
	iterateOverLinks(url, 2)

	var keys []string
	for k := range linkNumbers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", linkNumbers[k])
	}

}

// Perform a modified iterative deepening search, given a web address and a max depth
func iterateOverLinks(url string, numTimes int) {

	linkNumbers[url] = 1

	masterStack := arraystack.New() // The list of elements we need to visit

	masterStack.Push(url) // Adding the first element to the stack

	for numTimes > 0 { // each numTimes iteration is another depth level
		numTimes--

		retSet := hashset.New()      // Form a temporary set to hold elements from our current search
		for masterStack.Size() > 0 { // Used to visit all of the elements in the stack
			currentUrl, err := masterStack.Pop() // Get the top element of the stack
			if err == false {                    // If the stack is empty break; redundent
				break
			}
			tempSet := getPageWords(currentUrl.(string)) // Get the urls from a given site
			toStore := tempSet.Values()                  // Get the string values of the urls
			for _, value := range toStore {
				retSet.Add(value) // add the string values of our current search
			}
		}

		// Get the values our the search results from every web page we visited on a particular level
		// We know that our linkNames will not be on the linkNumbers map because we already check for this in getPageWords()
		linkNames := retSet.Values()
		for _, value := range linkNames {
			masterStack.Push(value)
		}
	}
}

// Adds http to urls that do not have it. This fixes an error of urls that dont have it.
func formatUrl(inputUrl string) string {
	if strings.HasPrefix(inputUrl, "https://") {
		return inputUrl
	} else if strings.HasPrefix(inputUrl, "http://") {
		return inputUrl
	} else {
		return ("http://" + inputUrl)
	}
	return "ERROR"
}

/*
Returns all of the links on a specified page

ie returns all of the links on the homepage of wikipedia

*/
func getPageWords(url string) *hashset.Set {
	returnSet := hashset.New()

	url = formatUrl(url) // Get the corrected url for a particular page

	response, err := http.Get(url) // Get the content of a web page
	// If there is an error, abandon ship!
	if err != nil {
		return returnSet
	}
	z := html.NewTokenizer(response.Body) // Get the text of a web page
	for {
		tt := z.Next() // Get each element of a web apge

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return returnSet
		case tt == html.StartTagToken:
			t := z.Token()
			for _, a := range t.Attr {
				// If an element is one we are looking for add it to the return set.
				if a.Key == "href" {

					var temp string
					// Store the name of the link
					temp = a.Val

					// Determine if we have seen this link before by seeing if it is in the map
					checkIfSeen := linkNumbers[temp]
					// If its not on the map (checkIfSeen == 0), add it and say we have seen it once
					// Else increment the number of times we have seen it on the map
					if checkIfSeen == 0 {

						if strings.HasPrefix(temp, "http") {
							linkNumbers[temp] = 1
							returnSet.Add(temp)
						}

					} else {
						numTimes := linkNumbers[temp]
						numTimes++
						linkNumbers[temp] = numTimes
					}
					break
				}
			}
		}
	}
	return returnSet
}

/*
	resp, _ := http.Get("http://www.lipsum.com/")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
*/
