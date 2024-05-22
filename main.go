package main

import (
    "fmt"
    "time"
)

func boilWater() {
    fmt.Println("Boiling water...")
    time.Sleep(2 * time.Second)
    fmt.Println("Water boiled")
}

func chopVegetables() {
    fmt.Println("Chopping vegetables...")
    time.Sleep(1 * time.Second)
    fmt.Println("Vegetables chopped")
}

func main() {
    go boilWater()
    go chopVegetables()

    time.Sleep(3 * time.Second)
    fmt.Println("Meal preparation complete")
}