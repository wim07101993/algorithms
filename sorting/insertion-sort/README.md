# Insertion Sort

Insertion sorting is one of the easiest ways to sort a list. It is not the most
performance but in contrast with most other sorting algorithms, it does not need
the entire list before it can start sorting.

The algorithm works by iterating over the items to sort and inserting each item
at the place where it is supposed to be according to the comparing method.

E.g.:

```
2, 5, 6, 3, 8, 4 ,7 ,9, 1
```

would sort like this:
```
2
2, 5
2, 5, 6
2, 3, 5, 6
2, 3, 5, 6, 8
2, 3, 4, 5, 6, 8
2, 3, 4, 5, 6, 7, 8
2, 3, 4, 5, 6, 7, 8, 9
1, 2, 3, 4, 5, 6, 7, 8, 9
```

![image](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)

The performance of this algorithm varies between $O(n)$ and $O(n^2)$ where
$O(n)$ is on an already sorted list and $O(n^2)$ on a reverse sorted list.