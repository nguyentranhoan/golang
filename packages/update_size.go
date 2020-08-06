package Media


func updateSize() []SizeImage {

	lines, err := ReadCsv("./config_size.csv")
	if err != nil {
		panic(err)
	}

	data := []SizeImage{}

	// Loop through lines & turn into object
	for _, line := range lines {
		data = append(data, SizeImage{
			width: line[0],
			height: line[1],
		})
	}

	return data
}
