package channel

// function takes the channel itself
// and put 1 into the channel
func Greet(ch chan int) {

	// this is how we put something to the channel
	// note that putting something in channel channel with face
	// the arrow head in front
	ch <- 1
}
