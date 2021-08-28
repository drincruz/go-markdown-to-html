A [trie](https://en.wikipedia.org/wiki/Trie), or prefix tree, is a fun little data structure. It's used in spell checkers and all sorts of auto-suggested input (search input, etc.). It is pretty useful.

## How it is Made

Well if we quickly look at how it works, we see that it is essentially just like any other tree; it contains nodes and those nodes have children. There is no limit to how many children a node can have like a binary search tree. This is because we will split each word and each node will have a character of the word. So, if we have "dog", "doggy", and "dogs" in our trie, the "g" node has three children nodes of "g", "s", and a termination node to represent an end of word, in this case the word, dog. Pretty neat eh?

If you still don't get it, something more visual usually helps out. This is Wikipedia's example.

![Trie example on Wikipedia](https://upload.wikimedia.org/wikipedia/commons/b/be/Trie_example.svg)

## Quite Simple Really

So take a peek at the [source](https://github.com/drincruz/python-trie) and see how it is really made. In Python we get to take advantage of nested dictionaries to hold our nodes. This is great since we can have a constant O(1) look up complexity at each character level.

So, going back to our simple example of "dog"; to find the word, we'll search the trie root for a "d". From there, we'll search the children nodes of "d" to see if there is an "o". Following that, we search for the "g". Then finally, we search for our termination node to confirm that "dog" is indeed in our trie. Simple right?

This is truly just a simple module at the moment, but obviously can be extended for a wide range of possibilities.
