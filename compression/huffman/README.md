# Huffman encoding

The huffman encoding is a variable length encoding. It compresses data by
storing the most used value using the least amount of bits.

To create a huffman encoding, a tree is needed which indicates which character
will be represented with which bit-string. This tree is created using a
huffman tree. Take for example the following string:

```
This is an example of the huffman tree.
```

1. Create a list of all the letters with their frequencies:

   | Letter | Frequency |
   |--------|-----------|
   |        | 7         |
   | .      | 1         |
   | ,      | 1         |
   | a      | 3         |
   | b      | 0         |
   | c      | 0         |
   | d      | 0         |
   | e      | 5         |
   | f      | 3         |
   | g      | 0         |
   | h      | 3         |
   | i      | 2         |
   | j      | 0         |
   | k      | 0         |
   | l      | 1         |
   | m      | 2         |
   | n      | 2         |
   | o      | 1         |
   | p      | 1         |
   | q      | 0         |
   | r      | 1         |
   | s      | 2         |
   | t      | 2         |
   | u      | 1         |
   | v      | 0         |
   | w      | 0         |
   | x      | 1         |
   | y      | 0         |
   | z      | 0         |
   | T      | 1         |

2. Order the list by frequency (and leave the 0 frequencies out)

   | Letter | Frequency |
   |--------|-----------|
   |        | 7         |
   | e      | 5         |
   | a      | 3         |
   | f      | 3         |
   | h      | 3         |
   | i      | 2         |
   | m      | 2         |
   | n      | 2         |
   | s      | 2         |
   | t      | 2         |
   | .      | 1         |
   | l      | 1         |
   | o      | 1         |
   | p      | 1         |
   | r      | 1         |
   | u      | 1         |
   | x      | 1         |
   | T      | 1         |

3. Create a binary tree

   Make a binary leaf of the least frequent letters and repeat.

   ```
   ' '7 e5 h3 f3 a3 i2 s2 m2 n2 t2 .1 T1 p1 l1 o1 u1 r1 x1
   
   ' '7 e5 h3 f3 a3 i2 s2 m2 n2 t2 .T2    pl2    ou2    rx2
                                   / \    / \    / \    / \
                                  .1 T1  p1 l1  o1 u1  r1 x2
   
   ' '7 e5 h3 f3  ai5    sm4    nt4     .Tpl4          ourx4
                  / \    / \    / \      /  \          /  \
                 a3 i2  s2 m2  n2 t2  .T2    pl2    ou2    rx2
                                      / \    / \    / \    / \
                                     .1 T1  p1 l1  o1 u1  r1 x2
   
   ' '7 e5  af6    ai5    sm4    nt4      .Tpl4          ourx4
            / \    / \    / \    / \       /  \          /  \
           h3 f3  a3 i2  s2 m2  n2 t2   .T2    pl2    ou2    rx2
                                        / \    / \    / \    / \
                                       .1 T1  p1 l1  o1 u1  r1 x2
   
   ' '7 e5  af6    ai5       smnt8             .Tplourx8
            / \    / \      /    \             /       \
           h3 f3  a3 i2   sm4    nt4      .Tpl4          ourx4
                          / \    / \       /  \          /  \
                         s2 m2  n2 t2   .T2    pl2    ou2    rx2
                                        / \    / \    / \    / \
                                       .1 T1  p1 l1  o1 u1  r1 x2
   
   ' '7  hf6    eai10        smnt8             .Tplourx8
         / \     / \        /    \             /       \
        h3 f3  e5  ai5    sm4    nt4      .Tpl4          ourx4
                   / \    / \    / \       /  \          /  \
                  a3 i2  s2 m2  n2 t2   .T2    pl2    ou2    rx2
                                        / \    / \    / \    / \
                                       .1 T1  p1 l1  o1 u1  r1 x2
   
     ' 'hf13    eai10        smnt8             .Tplourx8
      /  \       / \        /    \             /       \
   ' '7  hf6   e5  ai5    sm4    nt4      .Tpl4          ourx4
         / \       / \    / \    / \       /  \          /  \
        h3 f3     a3 i2  s2 m2  n2 t2   .T2    pl2    ou2    rx2
                                        / \    / \    / \    / \
                                       .1 T1  p1 l1  o1 u1  r1 x2
   
     ' 'hf13    eai10              mnst.lopruxT16
      /  \       / \               /           \
   ' '7  hf6   e5  ai5        smnt8             .Tplourx8
         / \       / \       /    \             /       \
        h3 f3     a3 i2    sm4    nt4      .Tpl4          ourx4
                          / \    / \       /  \          /  \
                         s2 m2  n2 t2   .T2    pl2    ou2    rx2
                                        / \    / \    / \    / \
                                       .1 T1  p1 l1  o1 u1  r1 x2
   
          ' 'hfeai23               mnst.lopruxT16
            /  \                   /           \
     ' 'hf13    eai10         smnt8             .Tplourx8
      /  \       / \         /    \             /       \
   ' '7  hf6   e5  ai5     sm4    nt4      .Tpl4          ourx4
         / \       / \    / \    / \       /  \          /  \
        h3 f3     a3 i2  s2 m2  n2 t2   .T2    pl2    ou2    rx2
                                        / \    / \    / \    / \
                                       .1 T1  p1 l1  o1 u1  r1 x2
   
               ' 'hfeaimnst.lopruxT39
                    /            \
         ' 'hfeai23               mnst.lopruxT16
           /  \                   /           \
    ' 'hf13    eai10         smnt8             .Tplourx8
     /  \       / \         /    \             /       \
   ' '7  hf6   e5  ai5     sm4    nt4      .Tpl4          ourx4
         / \       / \    / \    / \       /  \          /  \
        h3 f3     a3 i2  s2 m2  n2 t2   .T2    pl2    ou2    rx2
                                        / \    / \    / \    / \
                                       .1 T1  p1 l1  o1 u1  r1 x2
   ```
   
   This leads to the following tree
 
   ```
         ' 'hfeaismnt.Tplourx
              /         \
       ' 'hfeai         smnt.Tplourx
        /    \          /         \
    ' 'hf    eai     smnt         .Tplourx
     /  \    / \     /   \        /      \
   ' '  hf   e ai   sm   nt    .Tpl      ourx
        / \    / \  / \  / \   /   \     /   \
        h f    a i  s m  n t  .T   pl   ou   rx
                              / \  / \  / \  / \
                              . T  p l  o u  r x
   ```
   
4. When encoding add a 0 when turning left in the tree and 1 when turning right.
   This results in the following tree with values.

   ```
                                     ' 'hfeaismnt.Tplourx
                         /                                          \
              ' 'hfeai=0                                              smnt.Tplourx=1
              /         \                                  /                                 \
      ' 'hf=00           eai=01                    smnt=10                                     .Tplourx=11
         /  \              / \                      /    \                               /                     \
   ' '=000  hf=001    e=010  ai=011          sm=100        nt=101               .Tpl=110                         ourx=111
             / \               / \            / \            / \              /         \                        /       \
        h=0010 f=0011     a=0110 i=0111  s=1000 m=1001  n=1010 t=1011  .T=1100           pl=1101          ou=1110         rx=1111
                                                                          / \              / \              / \              / \
                                                                    .=11000 T=11001  p=11000 l=11011  o=11100 u=11101  r=11110 x=11111
   ```
   
   This results for our text in:
   
   ```
   T     h    i    s        i    s        a    n        e   x     a    m    p     l     e       o     f        t    h    e       h    u     f    f    m    a    n        t    r     e   e   .
   11000 0010 0111 1000 000 0111 1000 000 0110 1010 000 010 11111 0110 1001 11000 11011 010 000 11100 0011 000 1011 0010 010 000 0010 11101 0011 0011 1001 0110 1010 000 1011 11110 010 010 11000
   ```

   ```
   11111011001111010000011110100000010100100001011110001010001101111001010000110000011000101101100100000110111010011001110000010100100010111110001001011000
   ```
   
   This results in 152 bits or 19 bytes. The decoded text exists as 39 bytes.

5. To decode the string just follow the same path in the tree.
   
   ```
   111110110011110100000111...
   rrrrrlrrllrrrrlrllllrlll
       T   h   i   s      i...
   ```