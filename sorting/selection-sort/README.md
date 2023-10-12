# Selection sort

Selection sort is an in place sorting algorithm. It iterates over the list and
swaps the current item with the next smallest item.

E.g.:
```
2, 5, 6, 3, 8, 4 ,7 ,9, 1
```

would sort like this:
```
2, 5, 6, 3, 8, 4 ,7 ,9, 1
1, 5, 6, 3, 8, 4 ,7 ,9, 2
1, 2, 6, 3, 8, 4 ,7 ,9, 5
1, 2, 3, 6, 8, 4 ,7 ,9, 5
1, 2, 3, 4, 8, 6 ,7 ,9, 5
1, 2, 3, 4, 5, 6 ,7 ,9, 8
1, 2, 3, 4, 5, 6 ,7 ,9, 8
1, 2, 3, 4, 5, 6 ,7 ,9, 8
1, 2, 3, 4, 5, 6, 7, 8, 9
```

![image](https://upload.wikimedia.org/wikipedia/commons/9/94/Selection-Sort-Animation.gif)

This algorithm uses a time of $O(n^2)$ to sort. Making it impractical to use 
with large lists.