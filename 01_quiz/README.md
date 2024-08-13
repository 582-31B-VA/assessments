# Quiz 1: Bank accounts

> Before next class \
> Weight: 5%

Create the business logic for a banking application. It should be
possible to create a new account, as well as withdraw and deposit money
from or to said account. Those actions should be implemented as methods
on structs of type `account`. In addition to a balance, an account
should have an automatically generated unique identifier as well as an
owner. Withdrawing more money than there are funds in the account should
produce an error. An account's balance should be stored in cents.

Make sure to document the functions and the types of your application,
as well as to test them in a separate file.

## Submission

The project must be submitted in a GitHub Classroom repository. To
create the repository, click [here][], and accept the assignment.

[here]: https://classroom.github.com/a/HRMh0tzh


## Assessment criteria [5]

-   Readability [1]

    -   code is free of unused variables and functions
    -   use of whitespace/indentation is tidy and consistent
    -   long lines (~80 chars) are split
    -   whitespace is used to visually support logical separation
    -   variable/function names are consistent and descriptive
    -   constants are used instead of hard-coded values
    -   comments are present where warranted, prose is correct and
        helpful
    -   inline comments are used sparingly where needed to decipher
        dense/complex lines

-   Language conventions [1]

    -   no unnecessary use of obscure constructs 
    -   standard language features are used appropriately

-   Program design [2]

    -   program flow is decomposed into manageable, logical pieces
    -   function interfaces are clean and well encapsulated
    -   appropriate algorithms are used, and coded cleanly
    -   common code is unified, not duplicated
    -   program is well tested

-   Data structures [1]

    -   data structures are appropriate
    -   no redundant storage/copying of data
    -   no global variables
