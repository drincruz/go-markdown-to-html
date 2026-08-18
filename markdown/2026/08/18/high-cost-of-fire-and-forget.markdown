Projects all have different requirements and constraints. As engineers, we weigh a lot of pros and cons in our decision-making. Sometimes we deliberately cut scope, and sometimes those decisions come back to haunt us.

Let's take a look at a project I worked on that was rolled out, failed, and was ultimately brought back up to proper engineering standards.

![Photo by Greg from Pexels](pexels-greg-1251451-2418664.jpg)

## Tight Deadlines, Minimal Engineering Resources

Some projects seemingly set the tone straight from the beginning. For this project, we wanted to add audit logging to our application. But, of course, we had a strict deadline, and the engineers assigned to the project had limited bandwidth.

### Audit Logs / Activity Logs

Audit logging is a security feature in many products that gives customers the ability to review access and activity-tracking events like logins, logouts, permission changes, and organization-level updates.

Our design was pragmatic and straightforward: define the events to log, build the business logic to capture those events, persist them in a datastore, and expose a public API so customers could fetch their logs.

One key architectural decision involved a classic pros-and-cons discussion with our product manager. We advocated for a bit more time to introduce a queue into our infrastructure stack. However, our infrastructure team had zero availability, and waiting on them would have pushed our delivery date back by a month or more. Regrettably, we agreed to cut scope and opted for a "fire-and-forget" strategy instead.

## Issues on the Horizon

Implementation was fast and simple, and work progressed steadily despite the compressed timeline. During alpha testing with early customers, the data flowed well, feedback was positive, and we quickly iterated on a few recommendations.

We moved forward with the full rollout, and the engineers on the project transitioned back to other workstreams. Everything seemed to be running smoothly—until a perfect storm of MongoDB issues hit. The feature had been live for a few weeks without incident when database performance suddenly degraded, leaving us to figure out whether our audit logging system was the cause or merely a bystander.

Several teams that had shipped new features over the preceding month audited their code. On our end, we found a few slow queries that were bogging down MongoDB and optimized them. However, overall latency issues across Mongo persisted.

Ultimately, the root cause was not our audit logging feature itself, but an existing strain on our MongoDB instances that had finally hit a tipping point.

Yet, during our investigation, we discovered something concerning: we had experienced data loss, with several audit events failing to log silently. That discovery gave us all the justification we needed to prioritize the queueing architecture we had originally wanted.

## Post-Rollout, Post-Incident

Once the broader database issues were resolved, we scoped the queueing work with full support from our infrastructure team. The implementation of this asynchronous layer was surgical: deploy a new queue and worker, update the application logic to enqueue audit events, and let the background worker persist them into OpenSearch.

With AWS SQS and a Dead Letter Queue (DLQ) in place, the team could finally breathe easier knowing that transient errors wouldn't result in lost audit data.

## We Failed to Push Back

Deadlines are an inevitable reality of software engineering. Looking back, we should have pushed back harder against the timeline to ensure we built a durable, resilient system from day one. That quiet intuition in the back of your mind asking, _"What could go wrong here?"_ is almost always worth listening to.

#### Meta

Photo by Greg from Pexels: https://www.pexels.com/photo/seashore-scenery-2418664/
