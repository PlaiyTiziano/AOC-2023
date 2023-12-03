mkdir -p "$HOME/personal/AOC-2023/day-$1/chapters"

touch "$HOME/personal/AOC-2023/day-$1/main.go"

touch "$HOME/personal/AOC-2023/day-$1/chapters/chapter1.go"
touch "$HOME/personal/AOC-2023/day-$1/chapters/chapter1_test.go"
touch "$HOME/personal/AOC-2023/day-$1/chapters/chapter2.go"
touch "$HOME/personal/AOC-2023/day-$1/chapters/chapter2_test.go"

# Fetch the input
curl -H "Cookie: session=`cat $HOME/personal/AOC-2023/cookie.txt`" \
	"https://adventofcode.com/2023/day/$1/input" \
	-o "$HOME/personal/AOC-2023/day-$1/input.txt"
