package actions

import (
	"firstRest/database"
	"firstRest/handlers/web/Svg"
	"fmt"
	"time"
)

const (
	halvingPeriod = 1460 //days
	startX        = 100

	graphSizeX = 800
	graphSizeY = 500

	pointsX = graphSizeX - startX // Number of points in the period
	pointsY = 500
)

func Build(pair string) (string, error) {
	textLine1 := Svg.BuildText((graphSizeX)/2, 20, 16, "Graph")

	line1 := Svg.BuildLine(startX, 30, 100, graphSizeY-2, "black")
	line2 := Svg.BuildLine(100, graphSizeY-2, graphSizeX, graphSizeY-2, "black")

	textMax := Svg.BuildText(50, 480, 14, "Period max.")
	textMin := Svg.BuildText(50, 60, 14, "Period min.")

	svg := Svg.BuildSvg(graphSizeX+200, graphSizeY, textLine1, line1, line2, textMax, textMin)

	halvings, err := database.GetHalvings()
	if err != nil {
		return "", err
	}

	var lastHalving time.Time
	for index, halving := range halvings {
		data, err := database.GetAverageData(pair, halving.Date, lastHalving)
		lastHalvingData := lastHalving
		lastHalving = halving.Date

		if len(data) == 0 {
			continue
		}

		color := getColor(index)
		if color == "" {
			continue
		}

		svg += getPeriodData(data, color, 0)
		svg += Svg.BuildText(graphSizeX+100, (index+1)*50, 14, color+": "+lastHalvingData.Format("2006-01-02")+" - "+halving.Date.Format("2006-01-02"))

		fmt.Printf("Вытащили для периода %d %d записей \n", index, len(data))
		if err != nil {
			return "", err
		}
	}

	svg += currentHalving(pair, lastHalving, getColor(index+2))
	svg += Svg.BuildText(
		graphSizeX+100, (index+2)*50, 14,
		getColor(index+2)+": "+lastHalving.Format("2006-01-02")+" - "+
			time.Now().Format("2006-01-02"))

	svg += "</svg>\n"
	return svg, err
}

func currentHalving(pair string, lastHalving time.Time, color string) string {
	data, _ := database.GetAverageData(pair, time.Now(), lastHalving)
	if len(data) == 0 {
		return ""
	}

	first, _ := time.Parse(time.RFC3339, data[0].Date)
	last, _ := time.Parse(time.RFC3339, data[len(data)-1].Date)

	daysDifference := first.Sub(last).Hours() / 24
	c := pointsX * (daysDifference / halvingPeriod)

	//TODO тут мин макс надо предполагать на основе предыдущих годиков
	return getPeriodData(data, color, int(c))
}

func getColor(index int) string {
	switch index {
	case 0:
		return "blue"
	case 1:
		return "red"
	case 2:
		return "pink"
	case 3:
		return "green"
	case 4:
		return "red"
	case 5:
		return "black"
	default:
		return ""
	}
}

func getPeriodData(records []database.CoveragePriceRecord, color string, pointsX1 int) string {
	periodMax, periodMin := getMax(records), getMin(records)
	delta := periodMax - periodMin // 100%

	pointsLocal := pointsX
	if pointsX1 != 0 {
		pointsLocal = pointsX1
	}

	periodSvg := ""

	var lastX, lastY int
	lenRecords := len(records)

	step := lenRecords / pointsLocal
	for i := 1; i <= pointsLocal; i++ {
		pointRecord := records[i*step]
		posX := startX + i
		posY := ((pointRecord.AveragePrice - periodMin) / delta) * pointsY
		periodSvg += getSvgElement(posX, posY, lastX, lastY, color)
		lastX = posX
		lastY = int(posY)
	}

	return periodSvg
}

func getSvgElement(posX int, posY float64, lastX int, lastY int, color string) string {
	var elementSvg string
	if lastX == 0 {
		elementSvg += Svg.BuildCircle(posX, int(posY), 1, color)
	} else {
		elementSvg += Svg.BuildLine(lastX, lastY, posX, int(posY), color)
	}
	return elementSvg
}

func getMin(records []database.CoveragePriceRecord) float64 {
	min1 := records[0].AveragePrice
	for _, record := range records {
		if record.AveragePrice < min1 {
			min1 = record.AveragePrice
		}
	}
	return min1
}

func getMax(records []database.CoveragePriceRecord) float64 {
	max1 := records[0].AveragePrice
	for _, record := range records {
		if record.AveragePrice > max1 {
			max1 = record.AveragePrice
		}
	}
	return max1
}
