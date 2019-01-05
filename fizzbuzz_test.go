package fizzbuzz_test

import (
	"testing"

	"github.com/TonPC64/fizzbuzz"
)

func TestFizzBuzzShouldSayOne(t *testing.T) {
	result := fizzbuzz.Say(1)
	expected := "1"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayTwo(t *testing.T) {
	result := fizzbuzz.Say(2)
	expected := "2"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFizz(t *testing.T) {
	result := fizzbuzz.Say(3)
	expected := "Fizz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFour(t *testing.T) {
	result := fizzbuzz.Say(4)
	expected := "4"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayBuzz(t *testing.T) {
	result := fizzbuzz.Say(5)
	expected := "Buzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFizzBuzz(t *testing.T) {
	result := fizzbuzz.Say(15)
	expected := "FizzBuzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
