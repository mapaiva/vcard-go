# Contributing to vcard-go

So you're interested in contributing with `vcard-go`, that's awesome! :tada:. Here in this document, you're going to find all resources to guide you along the way from having an idea up to get your code merged.

## Bug reports

If you found any bug related to `vcard-go` please, before filing a new issue, consider searching for the problem using the issue tracker. If you still couldn't find a proper solution, file a new issue describing the problem with as many details you can. Here is a great issue template:

```
<Short description about the problem>

<Short tutorial describing how to reproduce the issue>

<Operating System specs (OS name, architecture, version)>

<go version>

<vcard-go version (release, commit hash)>
```

## Feature request

Open an issue exposing your thoughts and the overall picture of the idea before diving into any code. This way, the idea could be better defined, tested and specified.

## Pull Requests

vcard-go uses the "fork and pull" model [described](https://help.github.com/articles/about-collaborative-development-models/) here, where contributors push changes to their personal fork and create pull requests to bring those changes into the source repository. Pull requests are made against vcard-go's master-branch, tested and reviewed. Once a change is approved to be merged.


### Style

In order to make all code consistent across the entire project, vcard-go uses standards defined both by the community and by golang's architects. We follow strictly the amazing [Effective Go](https://golang.org/doc/effective_go.html) and the also amazing [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) to guide all codebase style definitions. Make sure your changes are according to these documents.

Moreover:

- Make sure you wrote tests for your changes
- Make sure your code is passing on `make test`
- Make sure your code is passing on `make lint`

### Commit Message Guidelines

We strongly recommend you to use Angular's [Commit Messages Guidelines](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-guidelines). We use it for generating release messages and Changelogs hence using it it's a must-have.
