[HyperLogLog in Practice: Algorithmic Engineering of a State of The Art Cardinality Estimation Algorithm](https://stefanheule.com/papers/edbt13-hyperloglog.pdf)

*Abstract*: Cardinality estimation has a wide range of applications and is of particular importance in database systems. Various algorithms have been proposed in the past, and the HyperLogLog algorithm is one of them. In this paper, we present a series of improvements to this algorithm that reduce its memory requirements and significantly increase its accuracy for an important range of cardinalities. We have implemented our proposed algorithm for a system at Google and evaluated it empirically, comparing it to the original HyperLogLog algorithm. Like HyperLogLog, our improved algorithm parallelizes perfectly and computes the cardinality estimate in a single pass.

*Labels*: #DataStructures
