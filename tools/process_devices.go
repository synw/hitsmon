package hitsmon

type Device struct {
	Props map[string]DeviceProps
}

type DeviceProps struct {
	IsBot          bool
	IsMobile       bool
	IsPc           bool
	IsTablet       bool
	IsTouchCapable bool
	UaString       string
}

func processUaString() {
	file, err := os.Open("devices.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened devices")
	defer file.Close()
}
