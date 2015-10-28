
package main
import (
    "fmt"
    "bufio"
    "os"

)

type player struct{
    xPosition int
    yPosition int
}

func parsePlayersFromStdIn() (player, player){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    var hero player
    var princess player
    yPosition := 0
    for scanner.Scan(){
        inputLine := scanner.Text()
        for xPosition,char := range(inputLine){
            if char == 112{
                hero = player{xPosition,yPosition}
            }
            if char == 109{
                princess = player{xPosition, yPosition}
            }
        }
        yPosition += 1
    }
    return hero, princess
}

func heroGetsPrincess(hero player, princess player){
    for {
        switch{

        }
    }
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    p1, p2 := parsePlayersFromStdIn()
    fmt.Println(p1)
    fmt.Println(p2)

}