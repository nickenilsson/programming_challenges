
package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

type player struct{
    xPosition int
    yPosition int
}

func parsePlayersFromStdIn() (player, player){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    gridSize, _ := strconv.Atoi(scanner.Text())
    var hero player
    var princess player
    yPosition := gridSize - 1
    for scanner.Scan(){
        inputLine := scanner.Text()
        for xPosition,char := range(inputLine){
            if char == 109{
                hero = player{xPosition,yPosition}
            }
            if char == 112{
                princess = player{xPosition, yPosition}
            }
        }
        yPosition -= 1
    }
    return hero, princess
}

func heroGetsPrincess(hero player, princess player){
    for {
        if hero.xPosition < princess.xPosition{
            hero.xPosition += 1
            fmt.Println("RIGHT")
        }
        if hero.xPosition > princess.xPosition{
            hero.xPosition -= 1
            fmt.Println("LEFT")
        }
        if hero.yPosition < princess.yPosition{
            hero.yPosition += 1
            fmt.Println("UP")
        }
        if hero.yPosition > princess.yPosition{
            hero.yPosition -= 1
            fmt.Println("DOWN")
        }
        if hero.xPosition == princess.xPosition && hero.yPosition == princess.yPosition{
            break
        }

    }
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    hero, princess := parsePlayersFromStdIn()
    heroGetsPrincess(hero, princess)
}