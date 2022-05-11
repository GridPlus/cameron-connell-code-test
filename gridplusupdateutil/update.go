package main

func main() {

	res, err := makeRequest()
	if err != nil {
		panic(err)
	}

	res.ProcessUpdates()

}
