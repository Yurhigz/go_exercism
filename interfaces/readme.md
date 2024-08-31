You will not write the code for the different languages yourself so you need to structure your code for the robot so that other developers can easily add more languages later.

As a first step, define an interface Greeter with two methods.

    LanguageName which returns the name of the language (a string) that the robot is supposed to greet the visitor in.
    Greet which accepts a visitor's name (a string) and returns a string with the greeting message in a specific language.

Next, implement a function SayHello that accepts the name of the visitor and anything that implements the Greeter interface as arguments and returns the desired greeting string. For example, imagine a German Greeter implementation for which LanguageName returns "German" and Greet returns "Hallo {name}!":

SayHello("Dietrich", germanGreeter)
// => "I can speak German: Hallo Dietrich!"

Now your job is to make the robot work for people that scan Italian passports.

For that, create a struct Italian and implement the two methods that are needed for the struct to fulfill the Greeter interface you set up in task 1. You can greet someone in Italian with "Ciao {name}!".

Before you call it a day, you are also supposed to finish the functionality to greet people in Portuguese.

For that, create a struct Portuguese and implement the two methods that are needed for the struct to fulfill the Greeter interface here as well. You can greet someone in Portuguese with "Olá {name}!".
