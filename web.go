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
	links := getPageWords(url)

	vals := links.Values()

	sites := arraystack.New()
	for i := 0; i < len(vals); i++ {
		sites.Push(vals[i])
	}

	for i := 0; i < len(vals); i++ {
		fmt.Println(sites.Pop())
	}

}

func getPageWords(url string) *hashset.Set {
	set := hashset.New()

	response, _ := http.Get(url)
	z := html.NewTokenizer(response.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return set
		case tt == html.StartTagToken:
			t := z.Token()
			for _, a := range t.Attr {
				if a.Key == "href" {
					//fmt.Println("Found href:", a.Val)
					var temp string
					temp = a.Val
					set.Add(temp)
					break
				}
			}
		}
	}

	return set
}

/*
	resp, _ := http.Get("http://www.lipsum.com/")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
*/
