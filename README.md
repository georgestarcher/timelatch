# timelatch [![Build Status](https://github.com/georgestarcher/timelatch/workflows/timelatch%20CI/badge.svg)](https://github.com/georgestarcher/timelatch/actions)[![Report Card](https://goreportcard.com/badge/github.com/georgestarcher/timelatch)](https://goreportcard.com/report/github.com/georgestarcher/timelatch)


A Go (golang) module for a simple time based latching bool flag.

Written by George Starcher

MIT license, check license.txt for more information
All text above must be included in any redistribution

## Installation

```shell
go get github.com/georgestarcher/timelatch
```

## Usage

```go

	package main

	import (
		"fmt"
		"github.com/georgestarcher/timelatch"
	)

	var testLatch TimeLatch
	testLatch.SetDefault()
	testLatch.LatchDuration = 1 * time.Second
	fmt.Printf"Test Latch: %v\n", testLatch)

	if !testLatch.IsLatched() {
		fmt.Println("Latching")
		testLatch.Timestamp = time.Now()
	} else {
		fmt.Println"Latched")
	}

	got := testLatch.IsLatched()
	expected := true

	if got != expected {
		fmt.Printf("IsLatched Active: Expected %v, got %v\n", expected, got)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf"Now: %v", time.Now())

	got = testLatch.IsLatched()
	expected = false

	if got != expected {
		fmt.Printf("IsLatched Expired: Expected %v, got %v", expected, got)
	}

```
