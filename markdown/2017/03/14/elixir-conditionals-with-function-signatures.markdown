## On Learning Elixir

So, I've been learning and doing a bit more [Elixir](http://elixir-lang.org/) lately. I'm only a couple months in, but Elixir has been the primary language I have been coding in at work. Happy days!

This is **not** an introductory write-up nor a tutorial, but rather a quick look at how conditionals and guards are done in Elixir. Onward.

## Let's Take a Look at FizzBuzz

So, yes, FizzBuzz; the classic programming problem and still a favourite interview question. As a quick refresher, this game is played by counting from one through _n_, but for multiples of three, we'll have "Fizz", multiples of five, "Buzz", and multiples of both, "FizzBuzz". I think that is the simplest I can word it.

Let's have a look at an example in Python first.

```python
def fizz_buzz(n):
    if 0 == (n % 3) and 0 == (n % 5):
        return "FizzBuzz"
    elif 0 == (n % 3):
        return "Fizz"
    elif 0 == (n % 5):
        return "Buzz"
    else:
        return str(n)
```

There is nothing crazy going on here, the logic is simple and it does exactly what we need it to. It uses the usual `if/else` logic that you would expect.

So, now let's look at a solution in Elixir.

```elixir
def fizz_buzz(n) when 0 === rem(n, 3) and 0 === rem(n, 5) do
  "FizzBuzz"
end
def fizz_buzz(n) when 0 === rem(n, 3), do: "Fizz"
def fizz_buzz(n) when 0 === rem(n, 5), do: "Buzz"
def fizz_buzz(n), do: n
```

If you've never done any Elixir before, I'm sure you may be a bit confused.

## Function Signatures as Conditionals

So, let me first say that, yes, `if/else` does exist in Elixir. "Then, why isn't it used here at all?", you ask? Well, Elixir has a good foundation of having their functions be explicit in what they do and being able to [pipe](http://elixir-lang.org/getting-started/enumerables-and-streams.html#the-pipe-operator) your code is very useful. I won't write much about the pipe operator here, but if you don't know much about it in Elixir, I highly recommend learning about it.

### Guards

So what we see in the Elixir FizzBuzz is an overloaded `fizz_buzz/1` function. As I'm sure you've looked at it and studied it a bit by now, yes, that is being done in lieu of `if/else`. The important piece is the when part of the function signature which is called a [guard](https://hexdocs.pm/elixir/master/guards.html). Using guards, we now have logic for which function we want to match on when we pass in the variable n. So now we see that each function signature matches up exactly to what we have done in Python with `if/elif/else` logic. Pretty neat eh?
