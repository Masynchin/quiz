# Game process

Game process consists of three steps:

- Connection to room.
- Solving quiz question.
- Viewing quiz resulsts.

## Game states

~~~mermaid
graph TD;
    Menu --> Room;
    Room --> Menu;
    Room --> Quiz;
    Quiz --> Results;
    Results --> Menu;
~~~

### Quiz

~~~mermaid
graph TD;
    ... --> Round;
    Round --> Round;
    Round --> ....;
~~~

### Round

~~~mermaid
graph TD;
    ... --> Question;
    Question --> ResultsYet;
    ResultsYet --> ....;
~~~

### Results

Page with quiz results.

## Full scheme

~~~mermaid
graph TD;
    Menu -->|"Search room"| Room;
    Room -->|"Quit room"| Menu;
    Room -->|"Owner starts game"| Round;
    Round -->|"Question timeout / All players answered"| RoundResults;
    RoundResults -->|"Next round"| Round;
    RoundResults -->|"All questions played"| GameResults;
~~~
