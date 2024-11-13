package main

/*

Go has tests.

Tests are functions that ensures the validity of your code.

Tests are part of packages, and are written in files with the _test.go suffix.
See examples in the internal/geometry package.

Tests are run with the go test command.

E.g.:
    go test ./...                                # Run all tests in the current module
    go test ./internal/geometry                  # Run all tests in the geometry package
    go test ./internal/geometry -run TestPointEq # Run the TestPointEq test in the geometry package

Additionally, it's possible to create a coverage report to see which parts of
your code are tested.

See script/test for an example of how to run tests and generate a coverage report.

A common pattern in Go tests is to create a structure for the tests so that you
can define all your cases in a concise way.

E.g.

    // file: number.go

    func BiggerThan10(n int) bool {
        return n > 10
    }

    // file: number_test.go

    func TestBiggerThan10(t *testing.T) {
        tests := []struct {
            n    int
            want bool
        }{
            {1, false},
            {10, false},
            {11, true},
        }

        for i, test := range tests {
            got := BiggerThan10(test.n)
            if got != test.want {
                t.Errorf("[%d] BiggerThan10(%d) = %v, want %v", i, test.n, got, test.want)
            }
        }
    }

    Here, we have an inline struct with the fields n and want. Then we define a
    list of those, and we loop through each one with the `range` operator.

    t.Errorf marks the test as failed and prints an error message.

Have a try at changing some logic in the geometry package and see how the tests fails.

Note: invalid syntax also fails the tests.

See https://pkg.go.dev/testing for more information on testing.

*/

// Intentionally left the main function out, it's not needed here.
// TODO: Write tests for the `Triangle` type.
// WIP
// TODO: Write tests for the `Person` type.
// DONE
// TODO: Write tests for the `People` type.
// WIP
