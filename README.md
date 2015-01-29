# Skewb

This is my second attempt at creating an optimal Skewb solver. It runs fairly fast and has found 11 move solutions in less than a second. It is slower than it could be because I do not multithread, but it's still acceptable.

This solver uses the standard move notation, but it takes input in an arbitrary way that I defined.

# 11 move scrambles

I created a program found in [difficult](difficult) which generates "difficult" scrambles&mdash;that is, 11 move scrambles. I like to use these to test the solver, although they might not be the most realistic tests. Some of these scrambles are as follows:

    B L B L B' U L' B U' B' L
    B L B L R B U' L' U' B' L
    B L B L R' B' L U B' R U'
    B L B L R' B' U R' U B' L
    B L B L R' U' B U' L U' B
    B L B L U L' R B' L R' B'
    B L B L U' B' R U' L B' U

There are 90 in total.