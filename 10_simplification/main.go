package main

/*

This file shows a couple of simpler way to do things you already know.

Some of the existing syntax is a bit verbose, and there are some shortcuts that
can be taken.

1. Variable declaration

    Currently, you can declare a variable like this:

    var x int = 10

    However, go can infer the type of the variable, so you can do this:

    x := 10

    It is important to note the :=, it declares and assigns a value to a
    variable at the same time.

    If you want to declare a variable without assigning a value, you can do
    this:
    var x int

    It is very common for instance to declare empty variables before a conditional assignment.

    E.g.

    var x int
    if someCondition {
        x = 10
    } else {
        x = 20
    }

2. For loop

    Currently, you can declare a for loop like this:
    var i int = 0
    for {
        if i > 10 {
            break
        }

        // Body of the loop

        i++
    }

    However, the usual form is to do:

    for i := 0; i <= 10; i++ {
        // Body of the loop
    }

    Here the initialization, condition and increment are all in the same line.
    And the variable i is only available in the scope of the loop.

    The main difference here is that in the long form, the condition was
    specified on when to break the loop. In the short form, the condition is
    specified on when to continue the loop.

    E.g.
        // loops from 0 to 10
        for i := 0; i <= 10; i++ {}

        // loops from 10 to 1
        for i := 10; i > 0; i-- {}

        // loops from 0 to 10 by 2
        for i := 0; i <= 10; i += 2 {}

    You can still use a break statement to stop the loop if needed.

    This will print all numbers from 0 to 4
    for i := 0; i <= 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println(i)
    }

    You can also use a continue statement to skip the rest of the loop and
    continue to the next iteration.

    This will print all numbers from 0 to 10 except 5.
    for i := 0; i <= 10; i ++ {
        if i == 5 {
            continue
        }
        fmt.Println(i)
    }


3. Functions return

    In a function, you can return a value more concisely like so:

    func add(a int, b int) int {
        return a + b
    }

    No need to create an intermediate variable to return the value.

4. Functions parameters

    If a function has multiple parameters of the same type, you can group them
    like so:

    // here a and b are of type int
    func add(a, b, c int) int {}

    // here a and b are of type int, and d and e are of type string
    func something(a, b, c int, d, e string) {}

5. Arrays declaration

    Currently, you can declare and assign an array like this:

    var a []int
    a = append(a, 0, 1, 2)

    However, you can declare and assign an array like this:

    a := []int{0, 1, 2}

6. Maps declaration

    Currently, you can declare and assign a map like this:

    var m map[string]int
    m = make(map[string]int)
    m["a"] = 1
    m["b"] = 2

    However, you can declare and assign a map like this:

    m := map[string]int{
        "a": 1,
        "b": 2,
    }
*/

func main() {
	// TODO: declare a variable `x` of type `int` and set it to 10
	// TODO: declare a variable `y` of type `int` and set it to 20
	// TODO: declare a variable `name` of type `string` and set it to "John"
	// TODO: declare a variable `z` of type `int` without assigning a value
	// TODO: assign the sum of `x` and `y` to `z`

	// TODO: write a loop that prints all numbers from 0 to 10
	// TODO: write a loop that prints all numbers from 10 to 0
	// TODO: write a loop that prints all even numbers from 0 to 10
	// TODO: write a loop that prints the first 10 odd numbers
	// TODO: write a loop that calculates the sum of the first 10 numbers
	// TODO: write a loop that calculates the sum of the first 10 odd numbers
	// TODO: write a loop that prints the first 10 numbers, except 5 and 7
	// TODO: write a loop that sums numbers until it reaches 20, print each new
	// sum
	// e.g. 0, 0+1, 0+1+2, 0+1+2+3, ...

	// TODO: write a function that sums 3 `int` and returns the result
	// TODO: write a function that multiplies 3 `int` and returns the result
	// TODO: write a function that returns the sum of all numbers between its
	// parameters.
	// If the second parameter is less than the first, return 0.
	// e.g. sum(1, 3) => 1 + 2 + 3 = 6
	//      sum(2, 5) => 2 + 3 + 4 + 5 = 14

	// TODO: declare an array of 5 integers and print it
	// TODO: declare an array of 5 strings and print it
	// TODO: declare a map with 3 key-value pairs and print it
}
