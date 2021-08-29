A very common question that I have been getting asked is, "Luigi? What's that?". Well, my answer that I usually give, in brief, is that it is a project open sourced by Spotify to facilitate workflow and dependencies. But to quote the [Luigi ReadTheDocs page](https://luigi.readthedocs.org/en/stable/):

> Luigi is a Python package that helps you build complex pipelines of batch jobs. It handles dependency resolution, workflow management, visualization, handling failures, command line integration, and much more.

## Luigi Tasks & Targets

The main pieces of Luigi are built around a `Task` and a `Target`. A `Task` is exactly what you would think it would be, it's a single task in your data pipeline. So for example, you may have a task that reads in some csv file and pull out specific values that you want from the file. The `Target`, is the intended output for your `Task`. So, back to our csv file example, an `output()` `Target` for that task may be a cleaned up csv file with the values you wanted.

## Here Is Our Example Task

```python
import luigi


class MyTask(luigi.Task):
    def output(self):
         return luigi.LocalTarget('my_output.csv')

    def run(self):
        with open('input.csv', 'r') as input:
            cat_count = 0
            for line in input.readlines():
                animal, age, color = line.split(',')
                if animal == 'cat':
                    cat_count += 1

        with self.output().open('w') as out_file:
             out_file.write(cat_count)
```

That is a [dumb] simple `Task`. This task has no requirements. All that it does is read in a file, parse it and write out its output to another file.

A few tings to note are the importance of `output()`. Luigi checks to see if the `output()` exists to check if this task is complete. That is Luigi's definition of complete. You can also override `complete()` if you do not have an output, but for now just think that every task needs an output.

## "So, what is this good for?"

So one big thing that I purposely left out when describing Luigi, is that it integrates really well into the whole Hadoop ecosystem! So, now let's take a step back and think about how we would process these batch data processing jobs without Luigi...

Let's say we have several jobs that need to be accomplished in order for the task you want to be considered complete. So for example, maybe we have a need to process some data that we have found in some log files. The log files are currently stored in S3, so we'll have a job to fetch those locally. Maybe, the logs need to be cleaned up a little bit, so we will do whatever filtering and et cetera transformations we'll need to do to cleanse the data. Next, we'll want the newly formatted data logs loaded into HDFS. After that, we can utilize Hive, so we'll want to create a table with those logs.

### Data jobs in the past

So in the past, we would just have these several jobs run in several `cron` jobs. But wait, depending on how much data we're working with, these jobs will have varying length of time to complete! So, the best you can do is see how long these jobs run and schedule them accordingly. So for example, we know the data is pulled down from S3 within ~20 minutes, so we'll schedule the next job 30 minutes later, and do the same thing for the remaining jobs as well.

### Now, with Luigi

With Luigi, you create dependency chains for jobs very easily by overriding the `requires()` method. Now when you define your entire process, you want it to run in this order: `Task0->Task1->Task2->Task3`. So, you can now say that `Task3` _requires_ `Task2` to run, `Task2` _requires_ `Task1` to run, etc. etc. This looks like the following:

```python
class Task3(luigi.Task):
    def requires(self):
        return [Task2()]

    """
    Other core code would go here as well, like run(), output()...

    """
```

Now, you have one single point to schedule and no need to guess when each job runs! Pretty neat right? :)

This is obviously just an introduction. I've only touched the surface about Luigi. But, if you need to build out data pipelines and enjoy doing so in Python, I highly recommend checking out Luigi! Cheers!
