Read in elements of a page, add all hrefs to a set (to ensure uniquity), add set to a stack and then iterate through the stack. During each iteration, pop the top of the stack and get all of the links from that page. If there is a link to the page we are looking for, end. Else, add all of those elements to a set. Continue adding all of the elements to a set until the stack is empty. At that point add all of the elements in our set to the stack and start over looking for the element we are looking for.

`See notes at bottom for more info`

Uses: https://github.com/emirpasic/gods

Requires

go get github.com/emirpasic/gods/lists/arraylist
go install github.com/emirpasic/gods/lists/arraylist

AND

go get golang.org/x/net/
go install golang.org/x/net/

`notes`

This is an attempt at iterative deepening as illustrated in this [gif](http://www.how2examples.com/artificial-intelligence/images/Iterative-Depth-First-Search.gif). I remove the possible recursion issue by adding items to a set after going to them, then checking the set before visiting again.

There is a bug that I discovered after downloading an old version of wikipedia and attempted to connect the pages for Hilary Clinton to Donald Trump and then did it backwards (this was 2007 wikipedia, so pre-election). The connection could be made 1 way but not the other, although I was able to so there is at least 1 bug.
