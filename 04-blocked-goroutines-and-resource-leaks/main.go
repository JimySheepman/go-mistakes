package main

type Result struct{}

type Search func(query string) Result

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

func Second(query string, replicas ...Search) Result {
	c := make(chan Result, len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

func Third(query string, replicas ...Search) Result {
	c := make(chan Result, 1)

	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		default:
		}
	}

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}
