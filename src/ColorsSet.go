package main

func getColorNameByRgb(rgb string, m map[string]string) string {
	return m[rgb]
}

func initColorsNamesSet() map[string]string {
	colors := make(map[string]string)

	colors["rgb(256,256,256)"] = "white"
	colors["rgb(256,0,0)"] = "red"
	colors["rgb(0,0,256)"] = "blue"

	return colors
}
