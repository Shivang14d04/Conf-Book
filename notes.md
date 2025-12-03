you have to google it to know which packages have which functionalities so that you cna use

# go run <nameOf File to run>

:= will only work for variables, not for constants and not if you wnat to specifically assign a type to a variable

we use package variables when we want to use them across multiple functions within the same package
so that we dont have to pass them as parameters to each function

we just need to write go in front of the function call to make it a goroutine
go routines run concurrently with other goroutines in the same program
we can use time.Sleep() function to delay the execution of a goroutine for a specified duration
