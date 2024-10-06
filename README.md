# Forward

Learning Go and basics of programming.

## Go

1. Install go: https://formulae.brew.sh/formula/go

2. Confirm it is installed: `go version`

    If you see something like `go version go1.XX.X <OS>/<ARCH>`, you are good to
    go.

3. Run the first file: `go run 00_setup/main.go`

    You can run any go file that start with `package main` by running `go run <path to file>`.

    For now, don't try to understand what `package main` means. We will get to
    that later.

## Exercises

Exercises are in different directories.

E.g. `00_setup` contains the first exercise.

Each exercise has a `main.go` file that you can run with `go run <path to
file>`.

To "solve" each exercise, look at the comments in the `main.go` file.
You will see `// TODO` and a task to complete.

Once completed, you can run the file and see the output.

There will be some lines/parts marked with `// IGNORE ME` that you should not
worry about. They are there to help with the making of the exercises, but are
not relevant to the exercise itself.

## Formatting your code

You can format your code by running `go fmt <path to file>`.

If you want to format all files in a directory, you can run `go fmt ./...`.

The PR will fail if your code is not formatted correctly, so make sure to run
it before submitting your work.

## Submitting your work

Before starting to work on a new exercise, create a new branch that will be used
only for this exercise.

E.g. `git checkout -b 00_setup`

Once you have completed the exercise, commit your changes and push them to
this branch. Then create a pull request to merge this branch. Any
review/comments will be added to the pull request.

Once everything looks good and the PR has been approved, merge it, delete the
branch and move on the next exercise.
