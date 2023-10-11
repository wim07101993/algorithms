# Selection sort

Selection sort uses two lists, a sorted and unsorted. It repeatedly searches
for the next smallest number in the unsorted list and adds it to the sorted
list. This algorithm is less performant than some other algorithms like 
insert-sorting but uses less write operations.

E.g.:
```
2, 5, 6, 3, 8, 4 ,7 ,9, 1
```

would sort like this:
```
2
2, 2
2, 2, 3
2, 2, 3, 4
2, 2, 3, 4, 5
2, 2, 3, 4, 5, 6
2, 2, 3, 4, 5, 6, 7
2, 2, 3, 4, 5, 6, 7, 8
1, 2, 3, 4, 5, 6, 7, 8, 9
```

This algorithm uses a time of $O(n^2)$ to sort. Making it impractical to use 
with large lists.