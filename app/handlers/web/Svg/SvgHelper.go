package Svg

import "fmt"

func BuildText(x int, y int, fontSize int, text string) string {
	return fmt.Sprintf(
		"<text x=\"%d\" y=\"%d\" font-size=\"%d\" text-anchor=\"middle\">%s</text>\n", x, y, fontSize, text)
}

func BuildLine(x int, y int, x2 int, y2 int, stroke string) string {
	return fmt.Sprintf(
		"<line x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" stroke=\"%s\" stroke-width=\"2\"/>\n", x, y, x2, y2, stroke)
}

func BuildCircle(cx int, cy int, r int, fill string) string {
	return fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"%d\" fill=%s />", cx, cy, r, fill)
}

func BuildSvg(graphHeight int, graphWidth int, text string, lineX string, lineY string, textMax string, textMin string) string {
	return fmt.Sprintf(
		"<svg width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\">\n%s%s%s%s%s",
		graphHeight, graphWidth, text, lineX, lineY, textMax, textMin,
	)
}
