package aws_assets

// func GetKey() {
// 	var data [][]string
// 	client := NewEc2Client()
// 	assets := sg.GetAssets(client)
// 	for _, asset := range *assets {
// 		if asset.Instance.KeyName != nil {
// 			fmt.Println(*asset.Instance.KeyName, asset.InstanceID, *&asset.Tags[0])
// 			data = append(data, []string{*asset.Instance.KeyName, asset.InstanceID, *&asset.Tags[0]})
// 		}

// 	}
// 	WriteToCSV(data, "./key.csv")

// }

// func WriteToCSV(data [][]string, filename string) {
// 	// 创建一个 CSV 文件
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// 创建一个 CSV writer
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// 逐行写入数据到 CSV 文件
// 	for _, record := range data {
// 		if err := writer.Write(record); err != nil {
// 			panic(err)
// 		}
// 	}
// }
