# Move all zeros in an array to its end

Write a method that moves all zeros in an array to its end. You should
maintain the order of all other elements. Here's an example:

    zerostoend([1, 0, 2, 0, 4, 0])
    // [1, 2, 4, 0, 0, 0]

Here's another one:

    zerosToEnd([1, 0, 2, 0, 4, 0])
    // [1, 2, 4, 0, 0, 0]

Can you do this without instantiating a new array?

## Analysis

I'm going to pick on this problem statement:

* Ha ha, both examples are the same!
* Using the word "method" when actually specifying a function seems louche.

Depending on the programming language chosen,
the candidate may not be able to modify an array in place with a function.
I don't think you can do that in Go,
but I used `[]int`, a "slice of integers",
which you can modify in a function.

[My code](a1.go) finds a zero-valued array element by
incrementing an index from zero.
Starting with the index one more than the index of a zero-valued
element, the code looks for a non-zero-valued array element.
It sets the array value of the zero-valued element to the non-zero-value,
and the value of the non-zero-valued element's index to zero.

This problem does let the job candidate show off knowledge
of arrays in that an array gets passed to a function,
and you have to know the difference between an index
and the value of the array at the index.
Any code that satisfies this problem almost certainly has to
keep at least 2 indexes into the array.

The hint/bonus about doing the problem without another array
is equivalent to asking about doing it in O(1) space.

Once again, it behooves the candidate to talk about test cases.
A single element, zero and non-zero would be good.
All zeros and all non-zeros would be good.
Several leading zeros with trailing non-zeros
and vice versa would be good.
Always try the example,
but be aware that these problem statements often have
apparently deliberately misleading examples.

### Around the web

* [Closest problem statement](https://algodaily.com/challenges/zeros-to-the-end)
* [Similar problem statement](https://www.geeksforgeeks.org/move-zeroes-end-array/)
* [A bit more elaborate](https://www.geeksforgeeks.org/move-all-zeros-to-start-and-ones-to-end-in-an-array-of-random-integers/)
