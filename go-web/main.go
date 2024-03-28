package main

import (
	"fmt"
	"syscall/js"
	"time"
)

type user struct {
	Name string
	Age  int
}

func add(this js.Value, inputs []js.Value) interface{} {
	fmt.Printf("add(%v, %v)\n", inputs[0], inputs[1])
	a := inputs[0].Int()
	b := inputs[1].Int()
	return a + b
  }

func getUserObj(this js.Value, inputs []js.Value) interface{} {
	fmt.Printf("getName(%v, %v)\n", inputs[0], inputs[1])

	return js.ValueOf(map[string]interface{}{
        "name":inputs[0].String(),
        "age": inputs[1].Int(),
    })
}

func returnBinaryData(this js.Value, inputs []js.Value) interface{} {
	// Example binary data (in real scenario, this could be generated or read from file etc.)
	binaryData := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21}

	// Create a Uint8Array from the binary data
	uint8Array := js.Global().Get("Uint8Array").New(len(binaryData))
	js.CopyBytesToJS(uint8Array, binaryData)

	return uint8Array
}

func calculatePi(this js.Value, inputs []js.Value) interface{} {
    iterations := inputs[0].Int()
    start := time.Now()

    pi := 0.0
    for i := 0; i < iterations; i++ {
        // The Leibniz formula alternates between addition and subtraction
        // depending on the term index (i).
        sign := 1.0
        if i%2 == 1 {
            sign = -1.0
        }

        // Calculate the current term
        term := 1.0 / float64(2*i+1)
        pi += sign * term
    }

    // Multiply by 4 to get the value of pi
    pi *= 4.0

    elapsed := time.Since(start)
    fmt.Println("Pi Calculation time Golang:", elapsed)

    return js.ValueOf(pi)
}

func main() {
  c := make(chan struct{}, 0)
  js.Global().Set("add", js.FuncOf(add))
  js.Global().Set("getUserObj", js.FuncOf(getUserObj))
  js.Global().Set("returnBinaryData", js.FuncOf(returnBinaryData))
  js.Global().Set("calculatePi", js.FuncOf(calculatePi))
  <-c
}