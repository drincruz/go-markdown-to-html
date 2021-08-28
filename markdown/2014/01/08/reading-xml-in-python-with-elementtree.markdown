```python
import xml.etree.ElementTree as ET
```

## Getting the XML Data Input

There are two ways I usually end up getting the XML data. One, opening the file itself if the XML is locally available to me:

```python
tree = ET.parse("filename.xml")
```

Or two, if the data is only available to me as a string:

```python
ET.fromstring(input_string)
```

## It Helps to Know the XML

In my examples, I am going reading an RSS 2.0 XML file. Similar to something like [this](https://cyber.law.harvard.edu/rss/examples/rss2sample.xml).


We see that the `rss` node is at the root, so we can get that like so:

```python
rss = tree.getroot()
```

If we wanted to iterate through the items in this RSS feed, we see that they are all children of the `channel` node. So we can get the channel with this:

```python
channel = rss.find('channel')
```

Now we can iterate in this channel for the items:

```python
for item in channel.iter('item'):
    print(item.find('title').text
```

## Supporting Namespaces

Yes, it's easy to support namespaces for easy parsing. All you need to do is pass a dictionary of the namespaces used, and you can continue your parsing fun.

```python
NAME_SPACES = { 'dc': 'http://purl.org/dc/elements/1.1/' }
```

So, to get a `dc:subject` node, you can do so like this:

```python
category = item.find('dc:subject', namespace=NAME_SPACES).text
```

## Simple, Yet Effective

There is obviously a lot more you can do, but this is just a really simple example!

Here is an entire example for reading xml as a file:

```python
import sys
import xml.etree.ElementTree as ET

def main():
    """
    Main
    """
    NAME_SPACES = {
        'dc': 'http://purl.org/dc/elements/1.1/'
    }
    root = ET.parse(sys.argv[1])
    rss = root.getroot()
    channel = rss.find('channel')
    for item in channel.iter('item'):
        title = item.find('title').text
        link = item.find('link').text
        description= item.find('description').text
        category = item.find('dc:subject', namespaces=NAME_SPACES).text
        print("%s\t%s\t%s\n\n%s\n\n\n" % (title,link,category,description))


if "__main__" == __name__:
    main()
```

Here is an entire example reading xml from a string (in this example, read from the web):

```python
import sys
import urllib
import xml.etree.ElementTree as ET

def main():
    """
    Main
    """
    NAME_SPACES = {
        'dc': 'http://purl.org/dc/elements/1.1/'
    }
    xml_feed = urllib.urlopen(sys.argv[1])
    root = ET.fromstring(xml_feed.read())
    channel = root.find('channel')
    for item in channel.iter('item'):
        title = item.find('title').text
        link = item.find('link').text
        description= item.find('description').text
        category = item.find('dc:subject', namespaces=NAME_SPACES).text
        print("%s\t%s\t%s\n\n%s\n\n\n" % (title,link,category,description))

if "__main__" == __name__:
    main()
```
