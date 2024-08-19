# Project: Twenty-one

> Due: September 3 \
> Weight: 20%

Twenty-one is the name given to a family of popular card games, the
progenitor of which is recorded in Spain in the early 17th century.
Whilst there are numerous variants of twenty-one, the following general
rules apply:

-   A *banker* deals two cards from the top of shuffled deck of 52 to a
    player known as the *punter*.
-   The punter, having examined both cards, announce whether they will
    stay with the cards they have ("stand") or receive another card from
    the banker ("hit").
-   The aim is to score exactly twenty-one points or, failing that, to
    come as close to twenty-one as possible. If a player exceeds
    twenty-one, they lose.
-   Number cards count as their number, the jack, queen, and king ("face
    cards" or "pictures") count as 10, and aces count as either 1 or 11.
-   Once the punter has either announced they will stay with their cards
    or exceeded twenty-one, the banker takes their turn.
-   The banker is dealt cards until their hand achieves a total of 17 or
    higher.
-   Whoever's hand is closer to twenty-one (but not higher) wins.

A standard 52-card deck comprises 13 ranks in each of the four suits:
clubs (♣), diamonds (♦), hearts (♥) and spades (♠). Each suit includes
three court cards (face cards): king, queen and jack. Each suit also
includes ten numeral cards, from one (ace) to ten.

## Requirements

Create a programs to play twenty-one from the command-line. When
executed, the program should print the punter's hand, and wait for them
to input "hit" to receive another card, or "stand" to stay with the
cards they have. Any other input should be responded to with an error
message explaining how to play. Once the punter has announced they will
stay with their cards or exceeded twenty-one, the program should print
the banker's hand, and declare a winner. At the end of a round, the
player is prompted to enter "continue" to play another round, or "exit"
to close the program. Of course, the deck must be shuffled before each
round. To do so, swap each card once with another random card in the
deck.

Functions and types should be documented using the idiomatic style, as
well as properly tested in a separate file.

## Tips

-   The punter only needs to know the total value of their hand, not
    which card they currently hold, or which card they have drawn.
-   You can refer to the [JavaScript implementation][] we did last semester
    in 21F.

[JavaScript implementation]:
https://github.com/582-21F-VA/assignments/blob/main/01_twenty-one/solution/twenty-one.js

## Submission

The project must be submitted in a repository using GitHub Classroom. To
create the repository, click [here][], and accept the assignment.

[here]: https://classroom.github.com/a/sHYQvqB6

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
    -   inline comments are used sparingly where needed to decipher
        dense/complex lines

-   Language conventions [2]

    -   no unnecessary use of obscure constructs 
    -   standard language features are used appropriately

-   Program design [10]

    -   program flow is decomposed into manageable, logical pieces
    -   function interfaces are clean and well encapsulated
    -   appropriate algorithms are used, and coded cleanly
    -   common code is unified, not duplicated
    -   program is well tested

-   Data structures [3]

    -   data structures are appropriate
    -   no redundant storage/copying of data
    -   no global variables
