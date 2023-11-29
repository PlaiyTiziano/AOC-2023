mkdir "$HOME/personal/AOC-2023/day-$1"

# Fetch the input
wget --header="Cookie: session=`cat $HOME/personal/AOC-2023/cookie.txt`" \
	-o "$HOME/personal/AOC-2023/day-$1/input.txt" \
	"https://adventofcode.com/2023/day/$1/input"
