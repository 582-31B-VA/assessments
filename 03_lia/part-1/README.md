# LIA: Part 1

> Due: September 17 \
> Weight: 20%

For this course's Learning Integration Assessment (LIA), you will build
a REST server for an online forum; i.e., a discussion site where people
can hold conversations in the form of posted messages. The assessment is
divided into three parts. The first part focuses on setting up the
application. You will find the requirements below.

## Requirements

-   The application should be organized as shown in class. Go files,
    HTML templates, and static ressources (style sheets, scripts, etc.)
    should be in separate directories. Helper functions, routes, and
    request handlers should be in separate files.

-   The server should have routes for the following pages: a homepage
    where users can browse discussion threads, a registration page where
    users can create an account, an account page where users can update
    their information, a page where users can create a new discussion
    thread, a page where users can read messages from a discussion
    thread, and a page where users can post a new message in a
    discussion thread.

-   Pages where users can submit data (see above) should have two
    routes: one for GET requests, and another for POST requests.

-   For now, pages won't have much content, but they should at least
    each contain an `<h1>` element with the title of the page.

-   Discussion threads and accounts will eventually be identified by a
    unique number. For those pages, the identifier should appear in the
    URL, and be available to the appropriate handler function.

-   GET requests with a URL starting with `/static` should be served the
    resource in the `static` directory whose path matches the rest of
    the URL.

-   Each page should have its own template which inherits from a base
    template. Parts of the web interface that are shared among templates
    should be split into partial templates. Although it's unclear yet
    what the website will look like, it's safe to assume it will have a
    header, a navigation bar, and a footer.

-   The server should use structured logging and dependency injection as
    shown in class. Each request should get logged. Log entries should
    include the request's method as well as its URI.

-   When launching the server in the terminal, it should be possible to
    use a flag to set the port the server listens to. The default port
    should be `4000`.

## Note

Everything needed to produce this assignment can be found in *Let's Go*
chapters 1, 2 and 3.

## Submission

The project must be submitted in a repository using GitHub Classroom. To
create the repository, click [here][], and accept the assignment.

[here]: https://classroom.github.com/a/Qr0SP1-i

## Assessment criteria [20]

-   Readability [5]

    -   code is free of unused variables and functions
    -   use of whitespace/indentation is tidy and consistent
    -   long lines (~80 chars) are split
    -   whitespace is used to visually support logical separation
    -   variable/function names are consistent and descriptive
    -   constants are used instead of hard-coded values
    -   comments are present where warranted, prose is correct and
        helpful
    -   interfaces are well documented
    -   inline comments are used sparingly where needed to decipher
        dense/complex lines

-   Language conventions [2]

    -   no unnecessary use of obscure constructs
    -   standard language features are used appropriately

-   Program design [10]

    -   requirements are met
    -   program flow is decomposed into manageable, logical pieces
    -   function interfaces are clean and well encapsulated
    -   appropriate algorithms are used, and coded cleanly
    -   common code is unified, not duplicated
    -   errors are handled correctly and provide context

-   Data structures [3]

    -   data structures are appropriate
    -   no redundant storage/copying of data
    -   no global variables

