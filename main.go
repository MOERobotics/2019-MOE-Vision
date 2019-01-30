package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
	"gonum.org/v1/gonum/stat"
)

// MainWindow Input Window
const MainWindow string = "MAIN WINDOW"

// GrayThresholdWindow Gray Threshold Window
var GrayThresholdWindow = "GRAY THRESHOLD WINDOW"

// HlsThresholdWindow HLS Threshold Window
var HlsThresholdWindow = "HLS THRESHOLD WINDOW"

// BgrThresholdWindow BGR Threshold Window
var BgrThresholdWindow = "BGR THRESHOLD WINDOW"

// ContourWindow Window with contours
var ContourWindow = "CONTOUR WINDOW"

// OutputWindow Out Window
var OutputWindow = "OUTPUT WINDOW"

// TrackBarHLS TrackBar to change HLS
var TrackBarHLS = "HSL TrackBar"

// TrackBarGray TrackBar to change gray threshold
var TrackBarGray = "Gray TrackBar"

// TrackBarB TrackBar to change blue
var TrackBarB = "B TrackBar"

// TrackBarG TrackBar to change green
var TrackBarG = "G TrackBar"

// TrackBarR TrackBar to change red
var TrackBarR = "R TrackBar"

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
	// start := time.Now()

	// srcFileName := os.Args[1]

	mainWindow := gocv.NewWindow(MainWindow)
	defer mainWindow.Close()

	grayThresholdWindow := gocv.NewWindow(GrayThresholdWindow)
	defer grayThresholdWindow.Close()

	hlsThresholdWindow := gocv.NewWindow(HlsThresholdWindow)
	defer hlsThresholdWindow.Close()

	bgrThresholdWindow := gocv.NewWindow(BgrThresholdWindow)
	defer bgrThresholdWindow.Close()

	contourWindow := gocv.NewWindow(ContourWindow)
	defer contourWindow.Close()

	outputWindow := gocv.NewWindow(OutputWindow)
	defer outputWindow.Close()

	trackBarHLS := hlsThresholdWindow.CreateTrackbar(TrackBarHLS, 255)
	trackBarHLS.SetPos(240)

	trackBarB := bgrThresholdWindow.CreateTrackbar(TrackBarB, 255)
	trackBarB.SetPos(245)

	trackBarG := bgrThresholdWindow.CreateTrackbar(TrackBarG, 255)
	trackBarG.SetPos(245)

	trackBarR := bgrThresholdWindow.CreateTrackbar(TrackBarR, 255)
	trackBarR.SetPos(245)

	trackBarGray := grayThresholdWindow.CreateTrackbar(TrackBarGray, 255)
	trackBarGray.SetPos(245)

	webcam, _ := gocv.OpenVideoCapture(0)
	defer webcam.Close()

	// webcam.Set(gocv.VideoCaptureFrameWidth, 320)
	// webcam.Set(gocv.VideoCaptureFrameHeight, 240)

	// srcImage := gocv.IMRead(srcFileName, gocv.IMReadColor)
	srcImage := gocv.NewMat()
	defer srcImage.Close()

	resizedSrcImage := gocv.NewMat()
	defer resizedSrcImage.Close()

	hlsImage := gocv.NewMat()
	defer hlsImage.Close()

	grayImage := gocv.NewMat()
	defer grayImage.Close()

	hlsImageFiltered := gocv.NewMat()
	defer hlsImageFiltered.Close()

	grayImageFiltered := gocv.NewMat()
	defer grayImageFiltered.Close()

	bgrImageFiltered := gocv.NewMat()
	defer bgrImageFiltered.Close()

	contoursImage := gocv.NewMat()
	defer contoursImage.Close()

	// fmt.Println(time.Since(start))

	for {
		// start = time.Now()
		webcam.Read(&srcImage)
		// fmt.Println("WEBCAM READ IMAGE", time.Since(start))

		// start = time.Now()
		gocv.Resize(srcImage, &resizedSrcImage, image.Point{}, 0.5, 0.5, gocv.InterpolationLinear)
		// fmt.Println("RESIZE", time.Since(start))

		mainWindow.IMShow(resizedSrcImage)

		// start = time.Now()
		gocv.CvtColor(resizedSrcImage, &hlsImage, gocv.ColorBGRToHLS)
		// fmt.Println("BGR TO HLS", time.Since(start))

		// start = time.Now()
		gocv.CvtColor(resizedSrcImage, &grayImage, gocv.ColorBGRToGray)
		// fmt.Println("BGR TO GRAY", time.Since(start))

		lowerHlsBound := gocv.Scalar{
			Val1: 0,
			Val2: float64(trackBarHLS.GetPos()),
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
			Val1: float64(trackBarB.GetPos()),
			Val2: float64(trackBarG.GetPos()),
			Val3: float64(trackBarR.GetPos()),
			Val4: 0,
		}
		upperBgrBound := gocv.Scalar{
			Val1: 255,
			Val2: 255,
			Val3: 255,
			Val4: 0,
		}

		gocv.InRangeWithScalar(hlsImage, lowerHlsBound, upperHlsBound, &hlsImageFiltered)
		gocv.InRangeWithScalar(resizedSrcImage, lowerBgrBound, upperBgrBound, &bgrImageFiltered)
		gocv.Threshold(grayImage, &grayImageFiltered, float32(trackBarGray.GetPos()), float32(255), gocv.ThresholdBinary)

		// hlsThresholdWindow.IMShow(hlsImageFiltered)
		// bgrThresholdWindow.IMShow(bgrImageFiltered)
		// grayThresholdWindow.IMShow(grayImageFiltered)

		// start = time.Now()
		contours := gocv.FindContours(grayImageFiltered, gocv.RetrievalExternal, gocv.ChainApproxSimple)
		// fmt.Println("FIND CONTOURS", time.Since(start))

		resizedSrcImage.CopyTo(&contoursImage)
		contourColor := color.RGBA{0, 0, 255, 0}
		centroidColor := color.RGBA{255, 0, 0, 0}

		// color0 := color.RGBA{255, 0, 0, 0}
		color1 := color.RGBA{0, 255, 0, 0}
		// color2 := color.RGBA{0, 0, 255, 0}
		// color3 := color.RGBA{0, 255, 255, 0}

		for index, contour := range contours {
			contourMoments := momentsFromContour(contour)

			if contourMoments["m00"] > 300 {
				gocv.DrawContours(&contoursImage, contours, index, contourColor, 1)

				contourCentroidX := contourMoments["m10"] / contourMoments["m00"]
				contourCentroidY := contourMoments["m01"] / contourMoments["m00"]

				centroidPoint := image.Point{
					int(contourCentroidX),
					int(contourCentroidY),
				}
				gocv.Circle(&contoursImage, centroidPoint, 1, centroidColor, -1)

				var lines [4]SlopeOffset

				// start = time.Now()
				for setIndex, points := range quadrilateralPoints(contour) {
					xs := make([]float64, len(points))
					ys := make([]float64, len(points))

					for subIndex, point := range points {
						// tempPoint := image.Point{
						// 	point.X,
						// 	point.Y,
						// }

						xs[subIndex] = float64(point.X)
						ys[subIndex] = float64(point.Y)

						// if setIndex == 0 {
						// 	gocv.Circle(&contoursImage, tempPoint, 1, color0, -1)
						// } else if setIndex == 1 {
						// 	gocv.Circle(&contoursImage, tempPoint, 1, color1, -1)
						// } else if setIndex == 2 {
						// 	gocv.Circle(&contoursImage, tempPoint, 1, color2, -1)
						// } else {
						// 	gocv.Circle(&contoursImage, tempPoint, 1, color3, -1)
						// }
					}

					slope, offset := stat.LinearRegression(xs, ys, nil, false)
					lines[setIndex] = SlopeOffset{
						offset,
						slope,
					}
				}
				// fmt.Println("FIND LINES", time.Since(start))

				gocv.Circle(&contoursImage, getIntersection(lines[0], lines[1]), 1, color1, -1)
				gocv.Circle(&contoursImage, getIntersection(lines[1], lines[2]), 1, color1, -1)
				gocv.Circle(&contoursImage, getIntersection(lines[2], lines[3]), 1, color1, -1)
				gocv.Circle(&contoursImage, getIntersection(lines[3], lines[0]), 1, color1, -1)
			}
		}

		contourWindow.IMShow(contoursImage)

		if mainWindow.WaitKey(33) >= 0 {
			break
		}
	}

}
