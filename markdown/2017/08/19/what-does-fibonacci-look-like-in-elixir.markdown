I recently had a conversation about Elixir since I have been using it more and it was not a language this person was familiar with. As curious engineers, one basic question that arose was, "what does Fibonacci look like?". I was happy to comply and provide some code.

As a quick refresher, we want to write a function that takes in an integer that represents the `nth` number in the Fibonacci sequence. For this implementation, we're assuming that the input is a non-negative integer. Okay, let's go!

## Recursive

So the typical recursive solution has a function `fib(n)` and we return an integer of 1 if `n` is equal to 0 or 1 and otherwise, we want to recursively call `fib` with the last two previous indexes.

```elixir
def fib(0), do: 1
def fib(1), do: 1
def fib(n), do: fib(n-1) + fib(n-2)
```

Happy days! We just need three lines to implement this! Now, if you recall my [previous post](/2017/03/14/elixir-conditionals-with-function-signatures.html) on Elixir function conditionals, that is the same thing we're doing here. We're returning 1 if we pass in a 0 or 1 index, otherwise we'll just do the recursive logic that we want. That's it!

## Iterative

"So what about an iterative solution?" you ask? Yes, that's actually what was discussed next anyhow. So, the typical solution for an iterative solution is to have a loop and have two variables track the `previous` and `current` values. So, that's fine, but something just felt wrong with that not being quite Elixir-y. So, I thought about it for a bit and came up with the following solution.

```elixir
def iter_fib(0), do: 1
def iter_fib(1), do: 1
def iter_fib(index) do
  Enum.reduce(2..index, [1, 1], fn(_i, acc) ->
    # Calculate the Fibonacci value
    fib_val =
      # Take the accumulator
      acc
      # Flatten the list since we're appending fib_val by a new list
      |> List.flatten()
      # Sum those values
      |> Enum.sum()
    [Enum.take(acc, -1), fib_val]
  end)
  # This is now the Fibonacci value for the index and its previous value, so just take the last value
  |> List.last()
end
```

So, what I've done here is create a list that represents the Fibonacci sequence. We have the same overloaded function signatures for index values of 0 and 1, and otherwise, the bulk of our code goes into our `Enum.reduce/3`. What we are doing is constantly keeping the list length at 2 so we can easily just sum the values and compute the next Fibonacci value. My first implementation actually was a bit memory hungry because I was just appending my list continuously and then taking a sum of the last two values. Why keep the list long if you only want to compute that last value right?

## And Bob's Your Uncle

That's it! Plain. Simple. Memory smart. Anyhow, you can checkout the full source [here in this gist](https://gist.github.com/drincruz/7700681371eff71da8da706cf998dc85). Cheers!

