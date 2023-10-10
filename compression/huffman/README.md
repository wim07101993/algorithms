# Huffman encoding

The huffman encoding is a variable length encoding. It compresses data by
storing the most used value using the least amount of bits.

To create a huffman encoding, a tree is needed which indicates which character
will be represented with which bit-string. This tree is created using a
huffman tree. Take for example the following string:

```
wooooooow this is short
```

1. Create a list of all the letters with their frequencies:

   | Letter | Frequency |
   |--------|-----------|
   | w      | 2         |
   | o      | 8         |
   |        | 3         |
   | t      | 2         |
   | h      | 2         |
   | i      | 2         |
   | s      | 3         |
   | r      | 1         |

2. Order the list by frequency (and leave the 0 frequencies out)

   | Letter | Frequency |
   |--------|-----------|
   | r      | 1         |
   | w      | 2         |
   | t      | 2         |
   | h      | 2         |
   | i      | 2         |
   |        | 3         |
   | s      | 3         |
   | o      | 8         |

3. Create a binary tree

   Make a binary leaf of the least frequent letters and repeat.

   ```
   r(1) w(2)
      \ /
     rw(3)

   r(1) w(2)  t(2) h(2)
      \ /        \ /
     rw(3)      th(4)
   
     r(1) w(2)  t(2) h(2)
        \ /        \ /
   i(2) rw(3)      th(4)
      \ /
     irw(6)
   
   t(2) h(2)    r(1) w(2)   (3) s(3)
      \ /          \ /        \ /
      th(4)   i(2) rw(3)       s(6)
                 \ /
                irw(5)
              
                r(1) w(2)                     
                   \ /                         
   t(2) h(2)   i(2) rw(3)    (3) s(3)
      \ /         \ /          \ /
     th(4)       irw(5)         s(6)
         \       /
          thirw(9)
              
                r(1) w(2)
                   \ /
   t(2) h(2)   i(2) rw(3)   (3) s(3)
      \ /         \ /         \ /
     th(4)       irw(5)        s(6) o(8)      
         \       /                \ /     
          thirw(9)                so(14)
              
                r(1) w(2)          
                   \ /              
   t(2) h(2)   i(2) rw(3)   (3) s(3)
      \ /         \ /         \ /
     th(4)       irw(5)        s(6) o(8)      
         \       /                \ /     
          thirw(9)                so(14)
                  \              /
                    thirw so(23)
   ```
   
   This leads to the following tree
 
   ```
          r w          
          \ /              
   t h   i rw    s
   \ /   \ /   \ /
   th    irw     s o      
    \    /       \ /     
     thirw       so
         \       /
         thirw so
   ```
   
4. When encoding add a 0 when turning left in the tree and 1 when turning right.
   This results in the following tree with values.

   ```
                    r(0110) w(0111)
                          \ /              
   t(000) h(001)  i(010) rw(011)   (100) s(101)
      \ /              \ /             \ /
     th(00)           irw(01)          s(10) o(11)
           \         /                     \ /     
            thirw(0)                    so(1)
                  \                  /
                         thirw so
   ```
   
   This results for our text in:
   
   ```
   w    o  o  o  o  o  o  o  w        t   h   i   s       i   s       s   h   o  r    t
   0111-11-11-11-11-11-11-11-0111-100-000-001-010-101-100-010-101-100-101-001-11-0110-000
   
   011111111111111111011110000000110101100010101100101001110110000
   ```
   
   This results in 63 bits or 8 bytes. The decoded text exists as 23 bytes.

5. To decode the string just follow the same path in the tree.
   
   ```
   0111111111111111110111100000001010101100010101100101001110110000
   w   o o o o o o o w      t  h  i  s     i  s     s  h  o r   t
   ```