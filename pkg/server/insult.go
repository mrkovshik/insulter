package insulter

import (
	"math/rand"
	"time"
	)


func pickWord(x []string) string {
	rand.Seed(time.Now().UnixNano())
	return x[rand.Intn(len(x))]
}

func Insult(name string) string {
	list := []string{
		"Handsome guy",
		"Fool",
		"Poop",
		"Mushroom",
		"Coon",
		"Winner",
		"Good guy",
		"Just a regular normal motherfucker",
	}
	swear:= pickWord(list)
	answer:= "Who is "+ name+ "today?"+name+"is a"+ swear+ "today!"
return answer
}
