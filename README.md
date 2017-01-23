# UBC Tech Talk

This repo is for the live coding portion of the talk. We will commit regularly to the `master` branch so you can follow along as we go.

## Project Set Up

You will require Go v1.7.4 and Node v5.5 to build the project on your local system.

## Interesting reads
[Why I Interned at a Startup Instead of a Tech Giant](https://medium.com/code-like-a-girl/why-i-interned-at-a-startup-instead-of-a-tech-giant-the-myth-of-the-good-job-170b8e54c7d5#.pdxdlvm5n)

[An opinionated guide to writing developer resumes in 2017](https://medium.com/@arthur_camara/how-to-write-a-good-resume-in-2017-b8ea9dfdd3b9?238743%3F1#.jbbqou5xr)

## About us
[Axiom Zen](https://axiomzen.co)

[ZenHub](https://zenhub.com)

## Tell us about you!
https://goo.gl/forms/lcPVfmXHT7aDTlC13

## Our Slides
[Download slides](https://github.com/axiomzen/ubc-tech-talk/files/720845/UBC.Tech.Talk.-.Jan.19.2017.pdf)

## Questions and Answers
> What other web stacks do you use?

- **Eric**:
  - For backend we've used mainly Ruby on Rails, Node with express or koa and now Golang.
  - For frontend we've used mostly Backbone, Angular, plain Javascript is more than sufficient too for a lot of MVPs. Now we're moving towards React + Redux.
  - For databases we've used Postgres a lot, MongoDB sometimes though we're finding that MongoDB doesn't hold well once a product becomes more complicated. We've also been using a lot of graph databases recently such as Blazegraph, OrientDB and Neo4j

> ZenHub for Safari?

- **Pablo**:
  - Implementing the Safari extension is not difficult, but publishing updates is. For Chrome and Firefox, this process can be automated and updates get propagated to the users in a couple of hours, for Safari, the update needs to be reviewed by Apple, and is not clear how this review is done, so we cannot guarantee our users that they will receive the latest features and bug fixes in a timely manner.

> What is your favourite meme

- **Chris**:
  - The original and best: Leeeeroy, Jeeeenkinnnns

> What kind of training will new interns receive?

- **Chris**:
  - Everything you need to know in order to provide a positive contribution

- **Eric**:
  - We do provide mentorship for interns but there's no better 'training' than actually jumping in and committing code.

- **Pablo**:
  - We have a good onboarding system. In ZenHub, we usually start by setting up the development environment and explaining how ZenHub works. We rely heavily on pull request reviews to mentor interns and to teach them good practices (or learn from them!), but we also do pair programming from if we need to, or put time aside for learning (follow tutorials, read documentation, teach about a specific pattern or technology). We also have check-ins every 2 weeks to align priorities and see if there is anything that we can improve to help our interns to learn and work more effectively. 

> which part of the code is the application of reflection?

- **Chris**:
  - I didn't really have time to dive into the details of this in the talk, but reflection is used in a number of places under the hood.  If you are interested, check out JSON [encoding](https://golang.org/src/encoding/json/encode.go) and [decoding](https://golang.org/src/encoding/json/decode.go) as well as one of the many places in the [database driver](https://github.com/go-pg/pg/blob/v5/types/scan_value.go).

> Annyeonghaseyo

- **Kan**:
  - Ne, annyeonghaseyo?

> Do you wanna build a snowman?

- **Chris**:
  - I'm so done with winter, Anna

- **Kan**:
  - Come on let's go out and play

- **Lucy**:
  - Already did this winter!

> Why didn't my question appear?

- **Chris**:
  - Try hitting refresh on your browser :)

- **Eric**:
  - We didn't have time to set up polling or socket.io so you'll have to refresh the browser to send another request to the API to get all questions again.

> high GPA vs projects?

- **Kan**:
  - Projects > GPA, GPA as tie-breaker

- **Eric**:
  - projects, hands down

- **Pablo**:
  - High GPA is a good measure of how well a student does in exams and homework, projects are evidence that the student can code and make things.

- **Lucy**:
  - Definitely projects, but GPA is a bonus for students and new grads. GPA is not always an indicator of success.

> Are you hiring?

- **Eric**:
  - Yes

- **Pablo**:
  - We are always looking for talented people.

- **Christine**:
  - YEAH!

- **Lucy**:
  - Yes, we're always on the look out for great talent!

> ?'); DROP TABLE questions;--,

- **Chris**:
  - The database driver sanitizes all input automatically - but nice try!

- **Kan**:
  - Little Bobby Tables, we call him

- **Pablo**:
  - Nice try

> Why not TypeScript?

- **Kan**
  - A few reasons. We've actually considered Typescript for a few of our longer standing projects, but moving a non-typed code base to a Typescript codebase was a lot of work. And when starting new projects, we figured if we wanted a language with static types it makes sense to choose a language where static types are built into the language, rather than as a patch

> How, aside from using go, did you reduce your Docker image size?

- **Chris**:
  - Part of it is our secret sauce - however we can tell you substantial savings can be had by moving to alpine-linux based containers.

> What's your favourite flavour of icecream?

- **Chris**:
  - Anything Häagen-Dazs®

- **Kan**:
  - Chocolate all the way

- **Eric**:
  - Strawberry

- **Pablo**:
  - Why choose one when you can have all of them?

- **Christine**:
  - Lemon sorbet

- **Lucy**:
  - Vegan or organic gelato

> What is your favourite breed of dog

- **Chris**:
  - Doge

- **Kan**:
  - American Eskimo

- **Eric**:
  - West Highland Terrier

- **Pablo**:
  - Hedgehogs

- **Christine**:
  - I'm scared of dogs

> Suggested resources to learn Go + React?

- **Eric**:
  - A tour of Go is great: https://tour.golang.org/welcome. For React there are a lot of tutorials out there like https://www.fullstackreact.com/30-days-of-react/

> Why does the display bug out when you overflow the display?

- **Kan**:
  - If you mean in the Editor, its because we zoomed in and didn't turn on the proper word wrap

> Thanks for the amazing demo!

- **Lucy**:
  - Our Pleasure!  Please fill out the survey and request the next topic for our talk!

> Any chance you guys will open source hash?

- **Chris**:
  - Open sourcing _Hatch_ is something we are strongly considering.  Stay tuned!

> I missed the deployment part. What was the deployment process like for this app?

- **Kan**:
  - [Heroku](https://devcenter.heroku.com/articles/deploying-go) with the Postgres Add-on. Heroku allows you to push a git repo to it, and it will handle most of the deployment for you. We did have to also server static files (the frontend)

- **Eric**:
  - For this particular app, we just used Heroku. Heroku now has an official buildpack for Golang and it is able to look up dependencies with govendor.

> What is the interview process like for full time positions

- **Eric**:
  - Our interview process is quite different than other companies but it is structured in a way to mimic closely what it would be like to be an Axiom. This benefits both of us since you'll be able to find out if we're the right fit for you.

- **Lucy**:
  - This will also be posted on our website shortly. Expect a dialogue with various team members and to do some hands-on coding. The process is slightly shorter for co-op roles.

> What's it like to work at ZenHub? What's your day-to-day life like?

- **Christine**:

  - Working at ZenHub is really fast paced and exciting. We ship a new extension about once a day so there's constantly new features to work on. When I started my internship on my first day I got assigned a bunch of bugs so I had to learn javascript and CSS and then after that I asked a lot of questions from my co-workers for help. On a typical day we usually have a big feature we are working on with a few other team members so we will be implementing that as well as reviewing other people's code and testing.
  
- **Lucy**:
  - Please read Christine's blog.

> can i haz cheezburger

- **Kan**:
  - u can haz

> What's it take to get an interview at Axiom Zen for a new grad?

- **Kan**:
  - Having personal projects help A LOT. Otherwise familiarity with, and keeping up with current web technologies help a ton too

- **Lucy**:
  - I'd like to add that we have interviewed people who didn't make it to the interview stage on their first or second. They come back to us when they gain more skills or experience. We're happy to see this so don't give up.

> Do you guys only make web apps, what about android or ios?

- **Chris**:
  - We we a ton of native iOS, with a sprinkling of Android and a dash of Unity.  We also are into VR and IOT!

- **Eric**:
  - A lot of native iOS, a little bit of Android, a little bit of Unity for VR

> What does your ideal inter / co-op applicant look like?

- **Kan**:
  - Someone who is a self starter and who is looking to learn, but able to work well independently without too much supervison, due to our desire to keep a distributed heirachy.

- **Eric**:
  - Someone receptive to feedback. We want all of our interns to be committing code on their first day in but just know that if you commit 10 lines of code, you may get 10 comments for your first PR review ;) And that's completely OK! That's how we all learn and improve.

- **Lucy**:
  - Strong communication and collaboration skills are important. We're also looking for people who can take feedback well and learn from mistakes because we're very honest with each other. We also want people who can hustle -- we're a startup after all!

> Who was your best co-op?

- **Christine**:
  - Me, of course :)

> What bear is best?

- **Chris**:
  - Black Bear of course...

- **Kan**:
  - Bears. Beets. Battlestar Galactica.
