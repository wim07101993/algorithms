# Binary search

Binary search is an algorithm that finds the position of a search value in a 
sorted list of items. It takes the value in the middle and compares it to the 
search value. If items are equal, the search is done. Else the search continues
only in the half in which the element can lie. 

In a linear searching algorithm the worst time it will take to find an item is
$O(n)$ where n is the number of elements in the list. Binary search reduces 
this time to $O(log(n))$. This is only an improvement when the list is not
small.

This algorithm only works on sorted lists.
 
![diagram from wikipedia](https://upload.wikimedia.org/wikipedia/commons/8/83/Binary_Search_Depiction.svg)
