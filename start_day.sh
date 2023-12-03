mkdir "$HOME/personal/AOC-2023/day-$1"

touch "$HOME/personal/AOC-2023/day-$1/main.go"

# Fetch the input
curl -H "Cookie: session=`cat $HOME/personal/AOC-2023/cookie.txt`" \
	"https://adventofcode.com/2023/day/$1/input" \
	-o "$HOME/personal/AOC-2023/day-$1/input.txt"
