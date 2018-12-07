# starkscript
starkscript provides a lexer for a my custom script language that was intended to use with my discord bot, StarkRavingMadBot.
Everything is hand-written in go. 
No clue if I'll ever implement it into the bot, but it's been a fun learning experience so far.
The design has been based off the [August 2011 talk](https://www.youtube.com/watch?v=HxaD_trXwRE) by Robert Pike on the subject,
and background I had during my final years at college.

The formal definition of the script is provided in the `langdef.ebnf` file, and an example script in `example.stark`

If you'd like to see the output of that example, most of my testing has been done running the `Main` test (`go test -run Main`).
