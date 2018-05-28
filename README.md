# ELAG 2018 Library Data APIs Bootcamp

[![Join the chat at https://gitter.im/elag18apis/Lobby](https://badges.gitter.im/elag18apis/Lobby.svg)](https://gitter.im/elag18apis/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Welcome to the open repository, documentation and materials for the ELAG 2018 Library Data APIs Bootcamp!

* When: Monday, June 4th, 2018, 10:00 - 16:00
* Where: [**National Library of Technology**](https://www.elag2018.org/how-to-access-ntk/)
* Workshop Materials: [github.com/cmh2166/elag18apis](https://github.com/cmh2166/elag18apis)
* Workshop Slides: to be linked / stored in this repo

## About the Bootcamp

A growing number of people are seeing the need to evolve library data systems to a data-forward microservices architecture. But, how to get there? This bootcamp gives an overview of a data API & service example, going from design to development to deployment. *It’s not meant to be an in-depth dive of every topic therein, but link these domains and topics together.* We hope that participants would then leave with more comfort on how to start separating out services and data needs for an evolution to clear RESTful APIs and scalable microservices.

[*free space where you can throw in any other buzzwords you'd like to see. Plus another reminder that we won't be doing a deep dive of each topic, but an accelerated walk through of an existing project and process from some recent work.*]

This bootcamp is a registration-only workshop as part of [ELAG 2018 in Prague](https://www.elag2018.org/). If you have questions about the workshop specifically, you can [contact the workshop leaders using the information below](#contact-before-during-after-the-bootcamp).

## Bootcamp Schedule

Time          | Topic                                           | Link(s)
------------- | ----------------------------------------------- | ------------------------------------------
10-10:30      | Bootcamp Introduction, Logistics, Goals         |
10:30-11      | Designing our API (ReST, PCDM, Swagger)         |
*11-11:10*    | *coffee break*                                  |
11:10-12:30   | Developing our API (Go, Go-Swagger, Localstack) |
12:30-13:30   | Lunch Break (on your own)                       |
13:30-14:20   | Deploying our API (Docker)                      |
*14:20-14:30* | *coffee break*                                  |
14:30-15:30   | Deploying our API (AWS)                         |
15:30-16      | Conclusion & Bootcamp Retrospective             |

## Our Expectations of You

### Social & Mental Prep

To keep this workshop a happy, safe and inclusive space, we ask that you review and follow [the Recurse Center Social Rules (aka Hacker School Rules)](https://www.recurse.com/manual#sub-sec-social-rules).

We also ask that you come willing to take part and help out where you may have expertise to share!

### Technical Prep

We request that all participants bring a laptop with internet connection & modern web browser. On that laptop, please already have:

- our [workshop GitHub repository](https://github.com/cmh2166/elag18apis) on your computer (with mechanism to update / pull down latest changes on Monday morning).
- installed the latest, stable version (18.03.1-ce) of [Docker Community Edition](https://www.docker.com/community-edition);
- installed latest, stable version (1.10.2) of [Go](https://golang.org/doc/install#install) & [set up your Go workspace](https://github.com/sul-dlss-labs/taco#go-local-development-setup).
- set up a [free AWS account](https://aws.amazon.com/free/) & [awscli](https://docs.aws.amazon.com/cli/latest/userguide/installing.html) connected to that free account.
- installed [localstack](https://github.com/localstack/localstack#installing) (this requires python, sorry) for use in your Go workspace.

**If you're short on time or other needs, you need at least Go, localstack, and docker for following along with local development.** We will use our lunch break to help folks catch up or debug these installation requirements as we are able, though the morning requires Go be set up.

If you have any issues with the above, please contact us ASAP using the [communication methods detailed below](#contact-before-during-after-the-bootcamp).

## Contact Before, During, After the Bootcamp

If you have questions or concerns leading up to or after the bootcamp, please user our Gitter channel to ask or open an issue on this GitHub repository, particularly with any questions dealing with workshop preparation or any installation issues. This allows multiple workshop leaders to respond as able, and other participants can also learn.
- Open an issue [here](https://github.com/cmh2166/elag18apis/issues);
- Join our Gitter channel for chatting [here](https://gitter.im/elag18apis/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link).

Both require that you login or create a free account with GitHub.

During the workshop, we will indicate the best ways to get help or communicate a question/comment - however, this bootcamp is intended to be informal, so feel free to speak up or indicate you have a question at any time.
