package alg

import (
	"math"
	"math/big"
	"math/rand"
	"reflect"
)

type RSA struct {
	// P, Q - простые числа, Fi = (P - 1) * (Q - 1)
	// N - длина алфавита (P * Q)
	// E - открытый ключ
	// D - закрытый ключ
	P, Q, Fi, N, E, D int64

	publicKey  PublicKey
	privateKey PrivateKey
}

func NewRSA(limit int64) *RSA {
	PrimeNumbers := GetPrimeNumbersFromN(limit)

	var p, q int64
	N := len(PrimeNumbers)
	rand.Intn(N)

	/*
		p = PrimeNumbers[len(PrimeNumbers)-1]
		q = PrimeNumbers[len(PrimeNumbers)-2]
	*/

	p = PrimeNumbers[rand.Intn(N)]
	q = PrimeNumbers[rand.Intn(N)]

	rsa := &RSA{P: p, Q: q, Fi: (p - 1) * (q - 1), N: p * q}

	PrimeNumbersToCreatePublicKey := GetPrimeNumbersFromN(rsa.Fi)
	for rsa.E == 0 {
		tmpValue := PrimeNumbersToCreatePublicKey[rand.Intn(len(PrimeNumbersToCreatePublicKey))]
		if ExtendedEuclideanGCD(tmpValue, rsa.Fi) == 1 {
			rsa.E = tmpValue
		}
	}
	/*

		for _, curPrimeNumber := range PrimeNumbersToCreatePublicKey {
			if ExtendedEuclideanGCD(curPrimeNumber, rsa.Fi) == 1 {
				rsa.E = curPrimeNumber
				break
			}
		}
		rsa.E = 65537
	*/
	rsa.D = OtherKeyByGCD(rsa.E, rsa.Fi)

	rsa.publicKey = PublicKey{N: rsa.N, E: rsa.E}
	rsa.privateKey = PrivateKey{N: rsa.N, D: rsa.D}

	return rsa
}

func (r *RSA) Encrypt(bytes []byte) []int64 {
	resultArray := make([]int64, 0)
	for _, b := range bytes {
		m := new(big.Int).SetBytes([]byte{b})
		c := new(big.Int).Exp(m, big.NewInt(r.E), big.NewInt(r.N))
		resultArray = append(resultArray, c.Int64())
	}
	return resultArray
}

func (r *RSA) Decrypt(encryptedData []int64) []byte {
	resultBytes := make([]byte, 0)
	var curByte []byte
	for _, num := range encryptedData {
		if num == 0 {
			curByte = []byte{0}
		} else {
			c := new(big.Int).SetInt64(num)
			m := new(big.Int).Exp(c, big.NewInt(r.D), big.NewInt(r.N))
			curByte = CreateByte(m)
		}
		resultBytes = append(resultBytes, curByte...)
	}
	return resultBytes
}

func CreateByte(m *big.Int) []byte {
	curByte := m.Bytes()
	if reflect.DeepEqual(curByte, nullArray) {
		curByte = append(curByte, byte(0))
	}
	return curByte
}

func GetPrimeNumbersFromN(n int64) (ps []int64) {
	ps = make([]int64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1)
	for i, l := int64(2), int64(math.Sqrt(float64(n))); i <= l; i++ {
		if !N[i] {
			for j := int64(2); i*j <= n; j++ {
				N[i*j] = true
			}
		}
	}

	for i, l := int64(2), n+1; i < l; i++ {
		if !N[i] {
			ps = append(ps, i)
		}
	}

	return ps
}

func OtherKeyByGCD(e, n int64) int64 {
	_, d, _ := ExtendedEuclideanGCDProcess(e, n)
	return d % n
}

func ExtendedEuclideanGCD(a, m int64) int64 {
	x, _, _ := ExtendedEuclideanGCDProcess(a, m)
	return x
}

func ExtendedEuclideanGCDProcess(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	} else {
		gcd, x, y := ExtendedEuclideanGCDProcess(b%a, a)
		return gcd, y - (b/a)*x, x
	}
}

func Int64ToBytes(array []int64) []byte {
	bytes := make([]byte, 0)
	for _, el := range array {
		big := new(big.Int).SetInt64(el)
		bytes = append(bytes, big.Bytes()...)
	}
	return bytes
}
