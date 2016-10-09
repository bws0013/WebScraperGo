package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
	"golang.org/x/net/html"
	"net/http"
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

	url := "http://auburn.edu/~bws0013/"

	//fmt.Println("Hello world")

	fmt.Println("stop 1")
	links := makeGlobalSet(url)

	vals := links.Values()

	sites := arraystack.New()
	fmt.Println("stop 2")
	for i := 0; i < len(vals); i++ {
		sites.Push(vals[i])
	}
	fmt.Println("stop 3")
	for i := 0; i < len(vals); i++ {
		fmt.Println(sites.Pop())
	}
	fmt.Println("stop 4")

}

/*
	Returns all of the links from all the pages a specified page links to

	ie returns all of the links on the all of the pages the homepage of wikipedia links to
*/
// TODO: Figure out why this throws an error.
func makeGlobalSet(url string) *hashset.Set {
	returnSet := hashset.New()

	addToReturnSet := getPageWords(url)

	vals := addToReturnSet.Values()
	fmt.Println("stop 6")
	for i := 0; i < len(vals); i++ {
		tempUrl := vals[i].(string)
		fmt.Println(tempUrl)
		tempAddToReturnSet := getPageWords(tempUrl)
		tempVals := tempAddToReturnSet.Values()
		for j := 0; j < len(tempVals); j++ {
			returnSet.Add(tempVals[j])
		}

	}
	fmt.Println("stop 5")
	return returnSet
}

/*
Returns all of the links on a specified page

ie returns all of the links on the homepage of wikipedia

*/
func getPageWords(url string) *hashset.Set {
	returnSet := hashset.New()

	response, _ := http.Get(url)
	z := html.NewTokenizer(response.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return returnSet
		case tt == html.StartTagToken:
			t := z.Token()
			for _, a := range t.Attr {
				if a.Key == "href" {
					fmt.Println("Found href:", a.Val)
					var temp string
					temp = a.Val
					returnSet.Add(temp)
					//bs.Add(temp)
					break
				}
			}
		}
	}
	fmt.Println("stop 7")
	return returnSet
}

/*
	resp, _ := http.Get("http://www.lipsum.com/")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
*/
