package main

import (
	"log"
	"os"
)

// printEnvs prints all environment variables
func printEnvs() {
	for _, v := range os.Environ() {
		log.Println(v)
	}
}

func main() {
	// print all environment variables
	log.Println("Environ:")
	printEnvs()

	// expand environment variables
	user := os.ExpandEnv("$USER, $HOME, $SHELL")
	log.Println("Expand USER, HOME, SHELL:")
	log.Println(user)

	// get value of environment variable
	user = os.Getenv("USER")
	log.Println("Getenv USER:")
	log.Println(user)

	if user, ok := os.LookupEnv("USER"); ok {
		log.Println("LookupEnv USER:")
		log.Println(user)
	}

	// clear all environment variables
	os.Clearenv()
	log.Println("Clearenv, Environ:")
	printEnvs()

	// set environment variable
	os.Setenv("USER", user)
	log.Println("Setenv USER, Environ:")
	printEnvs()

	// unset environment variable
	os.Unsetenv("USER")
	log.Println("Unsetenv USER, LookupEnv USER:")
	if user, ok := os.LookupEnv("USER"); ok {
		log.Println("User:")
		log.Println(user)
	}
}
