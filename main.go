package main

func main() {
	cfg := &config{
		commands: getCommands(),
	}

	startREPL(cfg)
}
