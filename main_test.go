package main

import (
	"testing"
	"math"
	"math/rand"
)

const epsilon = 1e-3 

func almostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func forwardKinematics(theta1, theta2, l1, l2 float64) (float64, float64) {
	x := l1*math.Cos(theta1*math.Pi/180) + l2*math.Cos((theta1+theta2)*math.Pi/180)
	y := l1*math.Sin(theta1*math.Pi/180) + l2*math.Sin((theta1+theta2)*math.Pi/180)
	return x, y
}

func TestInverseKinematics(t *testing.T) {
	tests := []struct {
		x, y   float64
		l1, l2 float64
	}{
		{6.0, 4.0, 5.3, 3.0},
		{6.0, 4.0, 6.0, 3.0},
		{6.0, 4.0, 4.0, 4.0},
		{9.0, 5.0, 5.3, 10.0},
		{0.001, 0.001, 5.3, 3.0},
		{5.3, 0.0, 5.3, 3.0},  
		{0.0, 5.3, -5.3, 3.0},  
		{6.0, 4.0, -5.3, 3.0},
		{6.0, 4.0, 5.3, -3.0},
		{6.0, 4.0, -5.3, -3.0},
		{9.0, 5.0, -5.3, 10.0},
		{5.3, 0.0, -5.3, 3.0},
	}

	t.Errorf("special error")
	for _, test := range tests {
			theta1, theta2, err := inverseKinematics(test.x, test.y, test.l1, test.l2)

			expectError := test.l1 < 0 || test.l2 < 0
			if expectError {
				if err == nil {
					t.Errorf("expected error for x=%.3f, y=%.3f, l1=%.3f, l2=%.3f, but got", test.x, test.y, test.l1, test.l2)
				}

			r := math.Sqrt(test.x*test.x + test.y*test.y)
			expectError = r > test.l1+test.l2 || r < math.Abs(test.l1-test.l2)

			if expectError {
				if err == nil {
					t.Errorf("expected error for x=%.3f, y=%.3f, l1=%.3f, l2=%.3f, but got", test.x, test.y, test.l1, test.l2)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for x=%.3f, y=%.3f, l1=%.3f, l2=%.3f: %v", test.x, test.y, test.l1, test.l2, err)
				} else {
					xf, yf := forwardKinematics(theta1, theta2, test.l1, test.l2)
					if !almostEqual(test.x, xf, epsilon) || !almostEqual(test.y, yf, epsilon) {
						t.Errorf("inverse kinematics failed for (%.3f, %.3f), got (%.3f, %.3f)", test.x, test.y, xf, yf)
					}
				}
			}
		}
	}
}

func TestRandomPoints(t *testing.T) {
	l1, l2 := 5.3, 3.0
	for i := 0; i < 1000; i++ {
		r := rand.Float64()*(l1+l2-math.Abs(l1-l2)) + math.Abs(l1-l2)
		angle := rand.Float64() * 2 * math.Pi
		x := r * math.Cos(angle)
		y := r * math.Sin(angle)

		theta1, theta2, err := inverseKinematics(x, y, l1, l2)

		if err != nil {
			t.Errorf("unexpected error for x=%.3f, y=%.3f, l1=%.3f, l2=%.3f: %v", x, y, l1, l2, err)
			continue
		}

		xf, yf := forwardKinematics(theta1, theta2, l1, l2)

		if !almostEqual(x, xf, epsilon) || !almostEqual(y, yf, epsilon) {
			t.Errorf("inverse kinematics failed for (%.3f, %.3f), got (%.3f, %.3f)", x, y, xf, yf)
		}
	}
}
