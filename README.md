# Samvera Connect 2018 Library Data APIs Bootcamp

Welcome to the open repository, documentation and materials for the Samvera Connect 2018 Library Data APIs Bootcamp!

* When: Tuesday, Otcober 9th, 2018, 9:00 - 17:00
* Where: [**J. Willard Marriott Library, University of Utah**](https://www.lib.utah.edu/info/directions.php)
* Workshop Materials: [github.com/PenguinParadigm/samvera18apis](https://github.com/PenguinParadigm/samvera18apis)
* Workshop Slides: [in Google Drive](https://docs.google.com/presentation/d/1wi5AlRt-r79xH-rQ1j6PbFTNz6bgqPVrznWt6HuNwYg/edit?usp=sharing)

## About the Bootcamp

A growing number of people are seeing the need to evolve library data systems to a data-forward microservices architecture. But, how to get there? This bootcamp gives an overview of a data API & service example, going from design to development to deployment. *It’s not meant to be an in-depth dive of every topic therein, but link these domains and topics together.* We hope that participants would then leave with more comfort on how to start separating out services and data needs for an evolution to clear RESTful APIs and scalable microservices.

[*free space where you can throw in any other buzzwords you'd like to see. Plus another reminder that we won't be doing a deep dive of each topic, but an accelerated walk through of an existing project and process from some recent work.*]

This bootcamp is a registration-only workshop as part of [Samvera Connect 2018](https://connect2018.lib.utah.edu/). If you have questions about the workshop specifically, you can [contact the workshop leaders using the information below](#contact-before-during-after-the-bootcamp).

## Bootcamp Schedule

Time          | Topic                                           | Link(s)
------------- | ----------------------------------------------- | ------------------------------------------
9-9:30        | Bootcamp Introduction, Logistics, Goals         | 
9:30-10:15    | Designing our API (ReST, PCDM, Swagger)         | 
*10:15-10:30* | *coffee break*                                  |
10:30-12      | Developing our API (Go, Go-Swagger, Localstack) |
12-14:00      | Lunch Break (on your own)                       |
14:00-14:45   | Deploying our API (Docker)                      |
*14:45-15*    | *coffee break*                                  |
15-16:30      | Deploying our API (AWS)                         |
16:30-17      | Conclusion & Bootcamp Retrospective             |

## Our Expectations of You

### Social & Mental Prep

To keep this workshop a happy, safe and inclusive space, we ask that you review and follow [the Recurse Center Social Rules (aka Hacker School Rules)](https://www.recurse.com/manual#sub-sec-social-rules).

We also ask that you come willing to take part and help out where you may have expertise to share!

### Technical Prep

We request that all participants bring a laptop with internet connection & modern web browser. On that laptop, please already have:

- our [workshop GitHub repository](https://github.com/PenguinParadigm/samvera18apis) on your computer (with mechanism to update / pull down latest changes on Monday morning).
- installed the latest, stable version (18.06.1-ce) of [Docker Community Edition](https://www.docker.com/community-edition);
- installed latest, stable version (1.11) of [Go](https://golang.org/doc/install#install) & [set up your Go workspace](https://github.com/sul-dlss-labs/taco#go-local-development-setup).
- set up a [free AWS account](https://aws.amazon.com/free/) & [awscli](https://docs.aws.amazon.com/cli/latest/userguide/installing.html) connected to that free account.
- pulled down & can run [this localstack docker image](https://hub.docker.com/r/localstack/localstack/) for use in your Go workspace.

**If you're short on time or other needs, you need at least Go, localstack, and docker for following along with local development.** We will use our lunch break to help folks catch up or debug these installation requirements as we are able, though the morning requires Go be set up.

If you have any issues with the above, please contact us ASAP using the [communication methods detailed below](#contact-before-during-after-the-bootcamp).

## Contact Before, During, After the Bootcamp

If you have questions or concerns leading up to or after the bootcamp, please user our Gitter channel to ask or open an issue on this GitHub repository, particularly with any questions dealing with workshop preparation or any installation issues. This allows multiple workshop leaders to respond as able, and other participants can also learn.
- Open an issue [here](https://github.com/PenguinParadigm/samvera18apis/issues);
- Join our #taco-workshop Samvera Slack channel.

During the workshop, we will indicate the best ways to get help or communicate a question/comment - however, this bootcamp is intended to be informal, so feel free to speak up or indicate you have a question at any time.
