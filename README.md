# Forward

Learning Go and basics of programming.

## Setup

1. Install go: https://formulae.brew.sh/formula/go

2. Confirm it is installed: `go version`

    If you see something like `go version go1.XX.X <OS>/<ARCH>`, you are good to
    go.

3. Run the first file: `go run 00_setup/main.go`

    It should fail with an error message. This is expected.

> [!NOTE]
> You can run any `main.go` file in this repo by running
> `go run <path to file>`

> [!IMPORTANT]
> Don't try setting up your editor with fancy plugins or anything yet. Just
> make sure you can run terminal commands and that your code is correctly
> highlighted. Any plugin will get in the way of learning the basics.

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

## Reading errors

If your code has an error, you will see an error message when running `go run
<path to file>`. This message will help you understand what is wrong with your
code.

E.g. if you run `go run 00_setup/main.go` without completing the exercise, you
will see something like:

```shell
$ go run 00_setup/main.go
# command-line-arguments // IGNORE ME
00_setup/main.go:19:10: syntax error: unexpected name code at end of statement
^ File path             ^ Error message
                 ^ Error line
                    ^ Error column
```

Pay close attention to any error messages you see, as they will help you
understand what is wrong with your code. If you don't understand the error
message, don't worry, just ask for help.

## Formatting your code

You can format your code by running `go fmt <path to file>`.

If you want to format all files in a directory, you can run `go fmt ./...`.

The PR will fail if your code is not formatted correctly, so make sure to run
it before submitting your work.

<details><summary>E.g.</summary>

In commit https://github.com/riraum/forward/commit/00bb7731ad7b4264bfdae135f475b659cb2ae4fd, I added a wrong indentation, and [the lint job failed](https://github.com/riraum/forward/actions/runs/11204189672/job/31142277128):

```shell
Error: running `go fmt ./...` results in modifications that you must check into version control:
diff --git a/00_setup/main.go b/00_setup/main.go
index bd232e6..a7f2504 100644
--- a/00_setup/main.go
+++ b/00_setup/main.go
@@ -23,3 +23,3 @@ func main() {
 	// E.g. (don't worry about the syntax yet):
-	  fmt.Println("Hello", "World!", 42, 3.14, true, false, []int{1, 2, 3})
+	fmt.Println("Hello", "World!", 42, 3.14, true, false, []int{1, 2, 3})
```

This means that the file `00_setup/main.go` is wrongly formatted, and should be
fixed.

</details>

## Submitting your work

Before starting to work on a new exercise, create a new branch that will be used
only for this exercise.

E.g. `git checkout -b 00_setup`

Make sure to commit often, ideally after each task is done.

Once you have completed the exercise, push your branch to GitHub. Then create a
pull request to merge this branch and request review from `nobe4`. Any
review/comments will be added to the pull request.

Once everything looks good and the PR has been approved, merge it, delete the
branch and move on the next exercise.
