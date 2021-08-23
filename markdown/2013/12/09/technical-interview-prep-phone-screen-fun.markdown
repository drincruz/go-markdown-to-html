So, I got some good feedback from my [Prep for your interview](/2013/11/15/prepare-for-your-interview.html) post. So, I've decided to cover some more topics in this realm. Hopefully, it will help a person or two. I won't be giving away the secret code to land any job, but I may have a hint or two.

## Technical Trivia / Useless Trivia

Let me first say, that I typically try and stay away from asking questions that need to be memorized. i.e. "What does the {insert_method_here} do in the {insert_random_library_here}?" Though, some interviewers like to ask them, so you better prepare just in case.

### Object oriented programming is usually a nice topic to be quizzed on over the phone.

#### What is a class? What is an object?

I like to think of a class as a blueprint of something. The object is the end result you get from the blueprint.

#### What is the difference between an interface and an abstract class?

Think of an interface as a contract. It will only describe what the implementing classes should look like. That being said, there are only signatures in an interface and an interface cannot do anything.
Now as for the abstract class, it actually is a class. It can define some methods to group similar class actions. For example, you can have a `Dog` abstract class with `walk()`, `bark()`, `eat()` as abstract methods. Then, specific dog types can extend this `Dog` abstract class and add their own individual methods that are specific to each class. So for example, maybe you can extend the class to be `YourDog`. Your dog is smart, it can `rollOver()`, `sit()`, `stay()` as well as the `walk()`, `bark()`, `eat()` actions. Silly example, but you hopefully get the point!

### Data structures

#### What is an array? What is a hashtable? Compare the two.

An array is a contiguous block of data that can be accessed via indices. A hashtable is a data structure that maps keys to values. It would be good to note that a hashtable is implemented with a hashing function to translate the keys into indices of a backing array to store the values.

## Paired Programming

If you thought coding while someone is watching you was weird, you may feel that similar distaste for paired programming on a phone screen. It's awkward. You have to try and find a solution to a problem while keeping the awkward silences to a minimum. That basically means, talking through how you are going to approach the problem. If you are anything like me, that means you'll end up saying something _completely_ nonsensical numerous times. To try and keep those to a minimum, try talking a little slower.

### FizzBuzz

I _love_ giving [FizzBuzz](https://en.wikipedia.org/wiki/Fizz_buzz) to folks. It should be relatively easy for you to get through this problem within a couple of minutes. Practicing coding problems is the way to go as they'll help you to be able to identify how to approach other problems.

## Lots to do, lots to review

Obviously, there is **a lot** of topics to cover. Interviewers all have their own style of interviewing. Prepare for the worst is always a safe bet and hopefully all will go fine!
