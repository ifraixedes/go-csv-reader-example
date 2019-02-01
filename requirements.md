# Requirements

This document contains the fictional requirements to be fulfilled.

## Description

Our marketing department would like some usage information regarding our cloud
Remote Builder service. We would like you to provide a very basic reporting implementation for
the data we have saved from the builds that have been executed.
You will be given a CSV file that contains information from our Remote Builder service that
executes builds in our cloud based system on user requests. The CSV file contents
will consist of fields in this order:

* Unique identifier for each build
* User ID reference the user that submitted the build request
* Time the build request was received (RFC 3339 formatted string)
* Time the build execution began (RFC 3339 formatted string)
* Time the build execution finished (RFC 3339 formatted string)
* Indicator for if the build has been deleted
* Exit code from the build process, >0 indicates failure
* Size of the resulting built image file

The marketing department wants to know how many builds were executed in a time window. For
example, how many builds were executed in the last 15 minutes, or in the last day, or between
January 1 and January 31, 2018.

The marketing department wants to know which users are using the remote build service the
most. Who are the top 5 users and how many builds have they executed in the time window?

The marketing department would like to know the build success rate, and for builds that are not
succeeding what are the top exit codes.

## Task

Design a simple solution to the situation described above using the Go programming language
and implemented as command-line program. The input will be a CSV file. The output should be
to stdout/stderr.

The implementation should include source code, tests, and a README.

The README should describe your object model, information about the software design choices
which has been made, and the testing approach.
