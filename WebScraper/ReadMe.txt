Read in elements of a page, add all hrefs to a set (to ensure uniquity), add set to a stack and then iterate through the stack. During each iteration, pop the top of the stack and get all of the links from that page. If there is a link to the page we are looking for, end. Else, add all of those elements to a set. Continue adding all of the elements to a set until the stack is empty. At that point add all of the elements in our set to the stack and start over looking for the element we are looking for.

Uses: https://github.com/emirpasic/gods

Requires 

go get github.com/emirpasic/gods/lists/arraylist
go install github.com/emirpasic/gods/lists/arraylist

AND 

go get golang.org/x/net/
go install golang.org/x/net/