# Delta encoding

The delta encoding is a lossless data compression algorithm. It stores the
difference between values in a list instead of the actual values.

The encoding is often combined with the RLE or huffman. After running the delta 
encoding, the can data contains a lot of the same runs of data.