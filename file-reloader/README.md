# File Reloader
An example of reloading a file (and updating its associated variables using a logical clock/variable's version number) as well as the test for performance benchmarks for different mechanisms to update/repopulate the variables storing the data (atomic pointers for the variables, atomic pointers for the version numbers, contexts with cancelations, mutexes).
