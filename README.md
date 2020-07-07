# Multiprocessing-with-golang
This little research carried out to know how efficeint the Golang concurrency feature is!
This little research carried out to know how efficient the Golang 'concurrency' feature is!
This script consists of the hybrid sorting algorithm (concepts of selection sort & mergeSort)
Any N sized array is divided into four (because I have 4 cores (2 real cores and another two are virtual core through hyper-threading in intel i3) sub-arrays and they are called concurrently for sorting.
Once these four sub-arrays are sorted, function merge merges those sub-arrays and returns a final sorted array.
below are the statistics for input size 1 million (10^6) integers,

1) simple selection sort in <b>Go</b> takes around <b>4 min</b> 
2) hybrid selection sort with <b>concurrency in Go</b> takes around <b>1 min</b>
3) simple selection sort in <b>C</b> takes around <b>3-4 min</b>
4) simple selection sort in <b>Python never gonna process</b> this sized input (it will take many days)

<b>note:</b> these results are machine specific and time complexity of selection sort will always remain O(N^2) for any programming language. but execution speed may vary from language to language.

<b>fun fact:</b>
if python's execution time is x then,
Go with concurrency (160x) faster > Go (40x) faster > C (10x) faster > python (x)
