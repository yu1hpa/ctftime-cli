package main

func main() {
	opts, err := ParseOpts()
	if err != nil {
		panic(err)
	}

	switch opts.Type {
	case "events":
		fetchEvents(opts)
	}
}
