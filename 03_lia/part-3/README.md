# LIA: Part 3

> Due: October 14 \
> Weight: 20%

For this course's Learning Integration Assessment (LIA), you will build
a REST server for an online forum; i.e., a discussion site where people
can hold conversations in the form of posted messages. The assessment is
divided into three parts. The third and finale part focuses on
processing forms and user authentication.

## Requirements

-   Pages for creating an account, creating a new discussion thread, and
    posting a message should now display a form allowing users to submit
    the necessary information.

-   Form data should be properly parsed when received by the server. If
    a field is invalid, the form should be displayed again with an error
    message indicating what the issue is. Except for passwords, fields
    should be automatically repopulated when there is an error.

-   No fields can be empty. Title of discussion threads cannot contain
    more than 100 characters. Messages cannot be more than 1000
    characters long. Email addresses need to be unique (use the method
    seen in class, not the one from the book) as well as sanity checked.
    Passwords need to be at least 8 characters long, and contain at
    least one uppercase character and one digit.

-   After an account, a discussion thread or a message has been created,
    users should be redirected to a page where they can see what was
    just created (i.e., the account page after creating an account, and
    the discussion thread page after creating a new thread or message).
    On this page, a flash message (or toast) should be displayed once,
    notifying users that the account/thread/message was successfully
    created.

-   Only logged in users can create discussion threads, post messages,
    and view their account. If a user tries to create a discussion
    thread or post a message without being logged in, they should get
    redirected to the login page. Also, the contents of the navigation
    bar should change depending on whether a user is authenticated or
    not : authenticated users should see links to "Home", "Create
    discussion thread" and "Logout", while unauthenticated users should
    see links to "Home", "Signup" and "Login".

## Note

Everything needed to produce this assignment can be found in *Let's Go*
chapters 7 (minus section 6), 8 and 10 (minus section 7).

## Submission

The project must be submitted in a repository using GitHub Classroom. To
create the repository, click [here][], and accept the assignment.

[here]: https://classroom.github.com/a/yRbTJQK3

Your first commit should only contain the code you submitted for part 2.
If you use the solution provided by the teacher, your first commit
should only contain the code in the solution. The message of the commit
should be `Add part 2 code`.

## Interview

Individual interviews will be conducted during the last class to assess
your understanding of the subjects covered throughout the semester. Your
ability to clearly explain your code will influence the mark you
receive.

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

