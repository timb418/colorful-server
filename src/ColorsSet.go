package main

func getColorNameByRgb(rgb string, m map[string]string) string {
	return m[rgb]
}

func initColorsNamesSet() map[string]string {
	colors := make(map[string]string)

	colors["rgb(256,256,256)"] = "white"
	colors["rgb(256,0,0)"] = "red"
	colors["rgb(0,0,256)"] = "blue"
	colors["rgb(220,20,60)"] = "crimson"
	colors["rgb(255,192,203)"] = "pink"
	colors["rgb(255,127,80)"] = "coral"
	colors["rgb(255,215,0)"] = "gold"
	colors["rgb(255,0,255)"] = "fuchsia"

	return colors
}
