package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]float64{{4.0, -1.0, 3}, 
							{1.0, 2.0, 4.0},
							{-3.0, 3.0, -5.0}}
	b := [][]float64{{5.0}, {0.0}, {-11.0}}

	
	fmt.Println(reverceMatrixMethod(matrix, b))
	
	fmt.Println(kramerMethod(matrix, b))
}

func kramerMethod(matrix, b [][]float64) [][]float64 {
	fmt.Println()
	fmt.Println("============Kramer's method=============")
	fmt.Printf("Matrix: \n")
	for i, val := range matrix {
		fmt.Print(val)
		fmt.Print("\t")
		fmt.Println(b[i])
	}
	fmt.Println()

	det := determinante(matrix)
	fmt.Printf("Determinante = %.4f\n", det)
	if det == 0 {
		fmt.Println("Error: det = 0")
		return nil
	}

	result := make([][]float64, 0)
	for i := 0; i < len(matrix); i+=1 {
		mi := kramer(matrix, b, i)
		fmt.Printf("Matrix for %d:\n", i)
		for _, val := range mi {
			fmt.Println(val)
		}

		dmi := determinante(mi)
		s := dmi/det
		fmt.Printf("Det for %d = %.4f/%.4f = %.4f\n\n", i, dmi, det, s)
		result = append(result, []float64{s})
	}

	fmt.Println("result:")
	return result
}

func kramer(matrix, b [][]float64, n int) [][]float64 {
	result := make([][]float64, 0)
	for i := 0; i < len(matrix); i+=1 {
		buff := make([]float64, 0)
		for j := 0; j < len(matrix[0]); j+=1 {
			if j != n {
				buff = append(buff, matrix[i][j])
			} else {
				buff = append(buff, b[i][0])
			}
		}
		result = append(result, buff)
	}
	return result
}

func reverceMatrixMethod(matrix, b [][]float64) [][]float64 {
	fmt.Println()
	fmt.Println("==========Reverce matrix method==========")
	fmt.Printf("Matrix: \n")
	for i, val := range matrix {
		fmt.Print(val)
		fmt.Print("\t")
		fmt.Println(b[i])
	}
	fmt.Println()

	det := determinante(matrix)
	fmt.Printf("Determinante = %.4f\n", det)
	algebraic_complement := make([][]float64, 0)

	if det == 0 {
		fmt.Println("Error: det = 0")
		return nil
	}

	for i := 0; i < len(matrix); i+=1 {
		buff := make([]float64, 0)
		for j := 0; j < len(matrix[0]); j+=1 {
			m := math.Pow(float64(-1), float64(i+j)) * determinante(getMinor(matrix, j, i))
			fmt.Printf("A(%d:%d) = %.2f\n", i, j, m)
			buff = append(buff, m)
		}
		algebraic_complement = append(algebraic_complement, buff)
	}
	fmt.Println()

	fmt.Printf("Reversive matrix:\n")
	algebraic_complement_transp := make([][]float64, 0)
	for i := 0; i < len(matrix); i+=1 {
		buff := make([]float64, 0)
		for j := 0; j < len(matrix[0]); j+=1 {
			buff = append(buff, algebraic_complement[j][i]/det)
		}
		fmt.Println(buff)
		algebraic_complement_transp = append(algebraic_complement_transp, buff)
	}
	
	fmt.Println()
	fmt.Println("result:")
	return matrixMultiplication(algebraic_complement_transp, b)
}

func matrixMultiplication(a, b [][]float64) [][]float64 {
	if len(a[0]) != len(b) {
		fmt.Println("Error: can not multiplicate this matrixs")
		return nil
	}
	
	result := make([][]float64, 0)
	for i := 0; i < len(a); i+=1 {
		buff := make([]float64, 0)
		for j := 0; j < len(b[0]); j+=1 {
			sum := 0.0
			for k := 0; k < len(a[0]); k+=1 {
				sum+=(a[i][k]*b[k][j])
				//fmt.Printf("A=%.3f B=%.3f\n", a[i][k], b[k][j])
			}
			buff = append(buff, sum)
		}
		result = append(result, buff)
	}

	return result
}

func getMinor(matrix [][]float64, x, y int) [][]float64 {
	result := make([][]float64, 0)
	for i, val := range matrix {
		buff := make([]float64, 0)
		for j, elem := range val {
			if j != x && i != y {
				buff = append(buff, elem)
			}
		}
		
		if len(buff) != 0 {
			result = append(result, buff)
		}
	}
	return result
}

func determinante(matrix [][]float64) float64 {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1]-matrix[1][0]*matrix[0][1]
	}

	det := 0.0
	for i := 0; i < len(matrix); i+=1 {
		if i % 2 == 0 {
			det += matrix[0][i] * determinante(minor(matrix, i))
		} else {
			det -= matrix[0][i] * determinante(minor(matrix, i))
		}
	}
	return det
}

func minor(matrix [][]float64, s int) [][]float64 {
	//var result [len(matrix)-1][len(matrix)-1]float64
	result := make([][]float64, 0)
	for i := 0; i < len(matrix)-1; i+=1 {
		buffArr := make([]float64, 0)
		for j := 0; j < len(matrix[0]); j+=1 {
			if j != s {
				buffArr = append(buffArr, matrix[i+1][j])
			}
		}
		result = append(result, buffArr)
	}
	return result;
}
