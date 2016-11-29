package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

/*
	First implementation:

	Read in elements of a page, add all hrefs to a set (to ensure uniquity), add set
	to a stack and then iterate through the stack. During each iteration, pop the top
	of the stack and get all of the links from that page. If there is a link to the
	page we are looking for, end. Else, add all of those elements to a set. Continue
	adding all of the elements to a set until the stack is empty. At that point add
	all of the elements in our set to the stack and start over looking for the element
	we are looking for.
*/
/*
	Improvements:
	Track the path of the bfs from the beginning to the end point. This could be done by
	mapping links the appended address of the links that led got it there. This would
	make the process slower.
*/

func main() {

	// url is the starting point of our search
	url := "http://auburn.edu/~bws0013/"
	//url := "https://www.facebook.com/"

	//links := getPageWords(url)

	// provide our url and the maximum depth we are trying to go within the web pages
	links := iterateOverLinks(url, 1)

	vals := links.Values()

	// All of the pages we visited in our search
	for _, value := range vals {
		fmt.Println(value)
	}

}

// Perform a modified iterative deepening search, given a web address and a max depth
func iterateOverLinks(url string, numTimes int) *hashset.Set {
	masterList := hashset.New()     // The list set ensuring we don't revisit web pages
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
		linkNames := retSet.Values()
		for _, value := range linkNames {
			if masterList.Contains(value) { // If we have visited the page before don't revisit
				continue
			} else {
				// If we have not visited the page add it as a page to visit and one not to visit again
				masterList.Add(value)
				masterStack.Push(value)
			}
		}
	}

	return masterList
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
					// fmt.Println("Found href:", a.Val)
					var temp string
					temp = a.Val
					if strings.HasPrefix(temp, "http") {
						returnSet.Add(temp)
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
