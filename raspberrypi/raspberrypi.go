package main

import (
	"fmt"
	"image"
	"image/color"
	"net"

	"gocv.io/x/gocv"
	"gonum.org/v1/gonum/stat"
)

func momentsFromContour(contourPoints []image.Point) map[string]float64 {
	tempMat := gocv.NewMatWithSize(len(contourPoints), 2, gocv.MatTypeCV32S)
	defer tempMat.Close()

	for index, point := range contourPoints {
		tempMat.SetIntAt(index, 0, int32(point.X))
		tempMat.SetIntAt(index, 1, int32(point.Y))
	}

	return gocv.Moments(tempMat, true)
}

func quadrilateralPoints(contourPoints []image.Point) [][]image.Point {
	quadPointsArray := make([][]image.Point, 0)

	isRestart := true
	startingIndex := -1
	currentSlopeSign := 0

	for index, point := range contourPoints {
		if isRestart && len(quadPointsArray) >= 3 {
			quadPointsArray = append(quadPointsArray, contourPoints[index:len(contourPoints)])
			break
		} else if isRestart {
			startingIndex = index
			skipAheadBy2Index := index + 2

			if skipAheadBy2Index < len(contourPoints) {
				deltaYBy2 := contourPoints[skipAheadBy2Index].Y - point.Y
				deltaXBy2 := contourPoints[skipAheadBy2Index].X - point.X

				if deltaXBy2 != 0 {
					slopeBy2 := float64(deltaYBy2) / float64(deltaXBy2)

					if slopeBy2 > 0 {
						currentSlopeSign = 1
					} else {
						currentSlopeSign = -1
					}

					isRestart = false
				} else {
					// log.Println("Divide by Zero", deltaYBy2, deltaXBy2, startingIndex)
					isRestart = true
				}
			}
		} else {
			skipAheadBy2Index := index + 2

			if skipAheadBy2Index < len(contourPoints) {
				deltaYBy2 := contourPoints[skipAheadBy2Index].Y - point.Y
				deltaXBy2 := contourPoints[skipAheadBy2Index].X - point.X

				if deltaXBy2 != 0 {
					slopeBy2 := float64(deltaYBy2) / float64(deltaXBy2)
					thisSlopeSign := 0
					if slopeBy2 > 0 {
						thisSlopeSign = 1
					} else {
						thisSlopeSign = -1
					}

					if thisSlopeSign != currentSlopeSign {
						// We have a change
						quadPointsArray = append(quadPointsArray, contourPoints[startingIndex:index+1])
						isRestart = true
					}
				}
			}
		}
	}

	return quadPointsArray
}

func getIntersection(line0 SlopeOffset, line1 SlopeOffset) image.Point {
	return image.Point{
		int((line1.offset - line0.offset) / (line0.slope - line1.slope)),
		int((line0.slope*line1.offset - line1.slope*line0.offset) / (line0.slope - line1.slope)),
	}
}

// SlopeOffset a struct for lines
type SlopeOffset struct {
	slope  float64
	offset float64
}

func main() {

	ips, _ := net.LookupIP("robot-2019.local")

	robotIP := []byte{}

	for _, item := range ips {
		if item.To4() != nil {
			robotIP = item.To4()
		} else {
			fmt.Println("NOT IPV")
		}
	}

	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: robotIP, Port: 5801, Zone: ""})
	defer Conn.Close()

	webcam, _ := gocv.OpenVideoCapture(0)
	defer webcam.Close()

	srcImage := gocv.NewMat()
	defer srcImage.Close()

	hlsImage := gocv.NewMat()
	defer hlsImage.Close()

	grayImage := gocv.NewMat()
	defer grayImage.Close()

	hlsImageFiltered := gocv.NewMat()
	defer hlsImageFiltered.Close()

	grayImageFiltered := gocv.NewMat()
	defer grayImageFiltered.Close()

	contoursImage := gocv.NewMat()
	defer contoursImage.Close()

	andImage := gocv.NewMat()
	defer andImage.Close()

	bgrImageFiltered := gocv.NewMat()
	defer bgrImageFiltered.Close()

	for {
		webcam.Read(&srcImage)

		gocv.CvtColor(srcImage, &hlsImage, gocv.ColorBGRToHLS)

		gocv.CvtColor(srcImage, &grayImage, gocv.ColorBGRToGray)

		lowerHlsBound := gocv.Scalar{
			Val1: 0,
			Val2: 0,
			Val3: 0,
			Val4: 0,
		}
		upperHlsBound := gocv.Scalar{
			Val1: 255,
			Val2: 255,
			Val3: 255,
			Val4: 0,
		}

		lowerBgrBound := gocv.Scalar{
			Val1: 0,
			Val2: 0,
			Val3: 0,
			Val4: 0,
		}
		upperBgrBound := gocv.Scalar{
			Val1: 255,
			Val2: 255,
			Val3: 255,
			Val4: 0,
		}

		gocv.InRangeWithScalar(hlsImage, lowerHlsBound, upperHlsBound, &hlsImageFiltered)
		gocv.InRangeWithScalar(srcImage, lowerBgrBound, upperBgrBound, &bgrImageFiltered)
		gocv.Threshold(grayImage, &grayImageFiltered, 0, float32(255), gocv.ThresholdBinary)

		contours := gocv.FindContours(grayImageFiltered, gocv.RetrievalExternal, gocv.ChainApproxSimple)

		color1 := color.RGBA{0, 255, 0, 0}

		for index, contour := range contours {
			contourMoments := momentsFromContour(contour)

			if contourMoments["m00"] > 300 {

				contourCentroidX := contourMoments["m10"] / contourMoments["m00"]
				contourCentroidY := contourMoments["m01"] / contourMoments["m00"]

				centroidPoint := image.Point{
					int(contourCentroidX),
					int(contourCentroidY),
				}

				Conn.Write([]byte(contourCentroidX))

				var lines [4]SlopeOffset

				for setIndex, points := range quadrilateralPoints(contour) {
					xs := make([]float64, len(points))
					ys := make([]float64, len(points))

					for subIndex, point := range points {

						xs[subIndex] = float64(point.X)
						ys[subIndex] = float64(point.Y)
					}

					slope, offset := stat.LinearRegression(xs, ys, nil, false)
					lines[setIndex] = SlopeOffset{
						offset,
						slope,
					}
				}
			}
		}
	}

}
