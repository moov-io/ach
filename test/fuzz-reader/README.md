## reader fuzzing

Fuzzing is a technique for sending arbitrary input into functions to see what happens. Typically this is done to weed out crashes/panics from software, or to detect bugs. In some cases all possible inputs are ran into the program (i.e. if a function accepts 16-bit integers it's trivial to try them all).

In this directory we are fuzzing `reader.go` from the root level -- In specific the `Reader.Read()` method.

### Running

If you need to setup `go-fuzz`, run `make install`.

Otherwise, run `make` and watch the output. Fix any crashes that happen, thanks!

See the `go-fuzz` project for more docs: https://github.com/dvyukov/go-fuzz
