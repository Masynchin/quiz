# Project models

There are categories of models:

- Game models
- Player models
- Quiz models
- Room models
- User models

## Quiz models

There are models of the quiz category:

- Quiz
- QuizQuestion
- Points
- QuizAnswer
- UserAnswer

### Quiz

Quiz - questions with the same topic.

Quiz has this protocol:

- Allow to get its text description.
- Allow to get its questions.

### QuizQuestion

QuizQuestion - question with options to answer it.

QuizQuestion has this protocol:

- Allow to get its text description.
- Allow to get options for answer.
- Allow to answer it with UserAnswer.

### Points

Points - score units for right answers the questions.

Points is just integer.

### QuizAnswer (internal model)

QuizAnswer - answer for QuizQuestion.

QuizAnswer has this protocol:

- Allow to compare it with UserAnswer.

### UserAnswer

UserAnswer - answer for QuizQuestion, given by User.

UserAnswer has this protocol:

- Allow to get its options.
