# GOTLifeTracker

This command line tool written in GO uses https://anapioficeandfire.com/ to
let you know if your favorite Game of Thrones character is dead or alive.

To use it, clone the repo then run "go build got_life_tracker.go"

Once you have the exectuable run "./got_life_tracker -firstName=$1 -lastName=$2"
where $1 is the first name of the character and $2 is the last name of the character

If you want to make it simpler to run from the command line, copy the executable to /usr/local/bin
Run "cp got_life_tracker /usr/local/bin"
Now you can just type "got_life_tracker -firstName=$1 -lastName=$2" without the "./"
