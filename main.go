package main

import (
	"fmt"
	"math"
)

func inverseKinematics(x, y, l1, l2 float64) (float64, float64, error) {

	if l1 <= 0 || l2 <= 0 { //
		return 0, 0, fmt.Errorf("lengths must be positive: l1=%.3f, l2=%.3f", l1, l2)
	}

	r := math.Sqrt(x*x + y*y) //

	if r > l1+l2 || r < math.Abs(l1-l2) {
		return 0, 0, fmt.Errorf("impossible to reach the point")
	} //

	cosA := (r*r - l2*l2 + l1*l1) / (2 * l1 * r)
	cosA = math.Max(-1, math.Min(1, cosA))
	angleA := math.Acos(cosA) //

	cosB := (-r*r + l1*l1 + l2*l2) / (2 * l1 * l2)
	cosB = math.Max(-1, math.Min(1, cosB))
	angleB := math.Acos(cosB) ////

	theta1 := (math.Atan2(y, x) - angleA) * 180 / math.Pi
	theta2 := (math.Pi - angleB) * 180 / math.Pi

	// [-180; 180]
	theta1 = normalizeAngle(theta1)
	theta2 = normalizeAngle(theta2)

	return theta1, theta2, nil
}

func normalizeAngle(angle float64) float64 {
	if angle > 180 {
		return angle - 360
	} else if angle < -180 {
		return angle + 360
	}
	return angle
}

func main() {
	x := 6.0
	y := 4.0

	l1 := 5.3
	l2 := 3.0

	theta1, theta2, err := inverseKinematics(x, y, l1, l2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("angle 1 %.2f degrees, angle 2 %.2f degrees\n", theta1, theta2) //
}
