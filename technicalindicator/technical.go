package technicalindicator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/markcheno/go-talib"
)

func Main() {
	fmt.Println("from technical!")
}

func SampleRSI() float64 {
	closings := []float64{
		44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08,
		45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64,
		46.21, 46.25, 45.71, 46.45, 45.78, 45.35, 44.03, 44.18, 44.22, 44.57,
		43.42, 42.66, 43.13,
	}

	return SampleRSIWithData(closings)
}

func SampleRSIWithData(closings []float64) float64 {
	values := talib.Rsi(closings, 14)
	if len(values) == 0 {
		return 0
	}

	return values[len(values)-1]
}

func StressTestSampleRSI(iterations, dataPoints int) (time.Duration, float64) {
	if iterations <= 0 || dataPoints <= 0 {
		return 0, 0
	}

	start := time.Now()
	last := 0.0
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < iterations; i++ {
		closings := make([]float64, dataPoints)
		price := 100.0 + rng.Float64()*10
		for j := 0; j < dataPoints; j++ {
			price += (rng.Float64() - 0.5) * 2
			if price <= 0 {
				price = 1
			}
			closings[j] = price
		}

		last = SampleRSIWithData(closings)
	}

	return time.Since(start), last
}
