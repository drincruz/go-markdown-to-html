A very common headache that I am sure every engineer has had to face at least one time in their life is character encoding. Oh, yes that fun topic! No, I do not have a solution for everyone. Sorry! But, in case you are in the Python world like I am and you are writing out files and getting a bunch of `UnicodeEncodeError`, well try the following below that I have.

## tl;dr Show me an example!

```python
import codecs

with codecs.open(filename, 'w', encoding='latin-1') as outfile:
    outfile.write('{}\n'.format(json.dumps(data, encoding='latin-1')))
```

## So, what? And why?

Okay, so I am writing a latin-1 encoded file. Yes, latin-1. Why? Well, I've chosen latin-1 here because latin-1 was giving me issues, so there! But really, if you want to write to a different encoding, obviously just swap that out.

But the long explanation, if you were curious, is that I am reading in data that was latin-1 encoded and it was making my data ingestion jobs fail because I default to utf-8. The `json.dumps()` bit is actually not needed if you are not working with json (obviously!). But, I wanted to point out that in case you were writing json, you also need to set the encoding to whatever you choose there as well. It is currently on my TODO list to see why that is the case.
