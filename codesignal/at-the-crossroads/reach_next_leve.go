package main

func main() {}

func reachNextLevel(experience int, threshold int, reward int) bool {
	return experience+reward >= threshold
}
