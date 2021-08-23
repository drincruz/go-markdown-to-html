So, I recently met a bunch of guys at a meetup. They were discussing the way their respective companies do interviews. I became intrigued, introduced myself, and we began to chat. For the majority, they either gave "homework" problems to solve or do some pair programming. So then, the fun part: **interview questions**.

For the most part, the questions were fairly easy; especially since some interview questions were so common that even I have used them when interviewing a candidate. But then, one of the guys wanted to test out an interview question and singled me out. He even wrote out on pen and paper the expected input and output of the program. I was really surprised! But most of all, I was now ridiculously **nervous**!

_The Problem_: Write a function that takes an input String array of names to pair up participants in "Secret Santa". The function should return a two-dimensional String array of the chosen pairs. Sounds easy, right! Yea, it should have been, but for whatever reason, my mind went blank.

I stumbled around with my words; probably said, "uh um" a dozen more times than I usually do. It was like I totally forgot how to program. It was **humiliating**!

The fellow was nice enough to give me hints as I finally started to find the right approach, but the solution was pretty sub-par. I hang my head but laugh it off. Me as an interview candidate: FAIL.

Now at home, that problem bugged me, so of course I took to an actual implementation. I finally was able to sit down and take my time with the problem. My issue was that I didn't attack the problem with the correct data structure. I was too busy fiddling with getting a random integer to get pairs, while I should have just shuffled the list, and used that list as a queue. Sheesh!

**TL;DR** When you are interviewing, try and relax, breathe; if you think you are over-thinking something, you probably are. Remember all of your data structures you have at hand!

Here is what I find a more acceptable solution than what I originally wrote out. So, to that fine gentleman that I met the other day, please excuse my poor coding the other night!

```java
import static java.util.Arrays.deepToString;
import java.util.Arrays;
import java.util.Collections;
import java.util.LinkedList;
import java.util.List;
import java.util.Random;
import java.util.Queue;

class SecretSanta {
    public String[][] go(String[] s) {
        List<String> partnerPool = new LinkedList<String>(Arrays.asList(s));
        Collections.shuffle(partnerPool);
        String[][] secretSantaList = new String[partnerPool.size()][2];
        for (int i=0; i<s.length; i++) {
            String partner = partnerPool.remove(0);
            if (s[i].equals(partner)) {
                partnerPool.add(partner);
                partner = partnerPool.remove(0);
            }
            secretSantaList[i][0] = s[i];
            secretSantaList[i][1] = partner;
        }
        return secretSantaList;
    }

    public static void main(String... args) {
        String[] players = new String[] {"Fred", "Wilma", "Barney", "Pebbles", "Bam Bam"};
        String[][] santas = new SecretSanta().go(players);
        System.out.println(deepToString(santas));
    }
}
```
