# Merge sort

Merge sort is a fast sorting algorithm in which two basic steps are involved:
- divide the list in sub-lists until each list contains only 1 element (when a 
  list only contains one element it is considered sorted)
- repeatedly merge the lists to produce merged lists until only one sorted list
  remains.

The image says it all:

![image from wikipedia](https://upload.wikimedia.org/wikipedia/commons/thumb/c/cc/Merge-sort-example-300px.gif/220px-Merge-sort-example-300px.gif
)

This algorithm has a worse speed complexity of $O(n)$ which makes it faster than
most other algorithms.