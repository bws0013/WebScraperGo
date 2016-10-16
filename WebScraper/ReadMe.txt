First implementation:

Read in elements of a page, add all hrefs to a set (to ensure uniquity), add set to a stack and then iterate through the stack. During each iteration, pop the top of the stack and get all of the links from that page. If there is a link to the page we are looking for, end. Else, add all of those elements to a set. Continue adding all of the elements to a set until the stack is empty. At that point add all of the elements in our set to the stack and start over looking for the element we are looking for.

Improvements:
	Track the path of the bfs from the beginning to the end point. This could be done by mapping links the appended address of the links that led got it there. This would make the process slower.