package main
import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    "strings"
)

type petrolPump struct{
    fuel int
    distanceNext int
}

func getPumpStations()[]petrolPump {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanner.Text()
    pumps := make([]petrolPump, 0)

    for scanner.Scan(){
        inputValues := strings.Split(scanner.Text()," ")
        fuel, _ := strconv.Atoi(inputValues[0])
        distanceNext,_ := strconv.Atoi(inputValues[1])
        pumps = append(pumps, petrolPump{fuel, distanceNext})
    }
    return pumps
}

func findStartingIndex(pumps []petrolPump) int {
    suitableStart := len(pumps)-1
    fuelDeficit := 0
    for i := len(pumps)-1; i >= 0; i-- {
        pumpFuelDelta := pumps[i].fuel - pumps[i].distanceNext
        if pumpFuelDelta + fuelDeficit >= 0{
            suitableStart = i
            fuelDeficit = 0
        }else{
            fuelDeficit += pumpFuelDelta
        }
    }
    return suitableStart
}

func main(){
    pumps := getPumpStations()
    startIndex := findStartingIndex(pumps)
    fmt.Println(startIndex)
}
