---
code: null
id: MDU6SXNzdWU0MTM3MzQxOTY=
databaseId: 413734196
number: 4
url: https://github.com/octolab/try/issues/4
title: "add streaming"
issueType: null
labels:
  - "type:feature"
  - "help wanted"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2019-02-23T20:07:38Z
updatedAt: 2019-02-26T06:53:22Z
lastEditedAt: 2019-02-23T20:08:00Z
closedAt: null
---

# add streaming

```go
retry --stream-to=12345 -- do command
```

retry starts listen 12345 TCP port to stream stdout/stderr of executed command to it.

can be useful https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

some inspiration:

* https://github.com/wagoodman/bashful
