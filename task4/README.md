## sre-gig-tasks
SRE Gig Q4 2023 Tasks

### Task4
Build a simple CLI app using Cobra package. This jokes-app is a CLI app that has the below subcommands:
- __get__: gets a joke from a remote Jokes API (https://icanhazdadjoke.com/) & stores the fetched joke in [jokes.txt](./jokes.txt) file
- __add__: adds a new joke from the CLI and stores it in the same local [jokes.txt](./jokes.txt) file
- __list__: lists all jokes stored in jokes.txt


#### Build jokes-app 
```
$ cd task4 && ls
cmd  go.mod  go.sum  main.go  README.md

$ go build && ls
cmd  go.mod  go.sum  jokes-app  main.go  README.md
```

#### Get a joke from https://icanhazdadjoke.com/
```
$ ./jokes-app get
My boss told me that he was going to fire the person with the worst posture. I have a hunch, it might be me.

$ ls 
cmd  go.mod  go.sum  jokes-app  jokes.txt  main.go  README.md

$ cat jokes.txt
My boss told me that he was going to fire the person with the worst posture. I have a hunch, it might be me.
```

### Steps to Run the Program

#### Add your own joke to [jokes.txt](jokes.txt)
```
$ ./jokes-app add 

Cannot add a null joke. Provide the joke as arguments to jokes-app add subcommand

$ ./jokes-app add Sore throats are a pain in the neck!
Adding joke "Sore throats are a pain in the neck!" 

$ cat jokes.txt
My boss told me that he was going to fire the person with the worst posture. I have a hunch, it might be me.
Sore throats are a pain in the neck!
```

#### list all jokes from [jokes.txt](jokes.txt)
```
$ ./jokes-app list

My boss told me that he was going to fire the person with the worst posture. I have a hunch, it might be me.
Sore throats are a pain in the neck!
```