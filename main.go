package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

var errInt int
var successInt int

func main() {
	kissing := `
	    	,,(%%%%%%%%%%%%%%%%%%%%%%%%%%%%#/  .     .....,,,***////////****,,,........................,,***////
		%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%#,..........,,,,**************,,,,,...............,,.,,,,,,,****
		%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%*.......,,,,,****************,,,,,,....... ........,,,,,,***
		%%%%%%%%%%%%%%#%%%#%%%%%%%%%%%%%%%%%%%%%%%%* ......,,,,**************,,,,,............. ......,,,,**
		%%%%%%%%%%##%%#%%###%#%%#%%%##%%%%%%#####%%%%(........,,,,,,,,*********,,,,.,....................,,,
		%%%%%%%%######%####%#############%#%%##((#((((%/.  .. .......,,,*****/**,/#%%%%%%%%%%%%%%%%%%%%#(/,.
		%%%%%%##%%#%%%#####(####((###(#########(##(#(((#,...............,,,,*#%%%####%%%%%%%%%%%%%%%%%%%%%%&
		%%##%##%###%%%%#########((#(#(((((#####((((###(#(* ..       .....,#(((((((((((#%%%%%%%%%%%%%%%%%%%%%
		%%%##############(##(#((####(((((((####((((#######(............*((((((((((((##%%%#%%#%%%%%%%%%%%%%%%
		%%%#######################(((((((((#((#((((########*,....... (((((((((((((((#%######%##%%%%%%%%%%%%%
		#%%%%#############################((((###(#(#######(,,,....,#((###(##((((((((((###########%%%%%&%%#%
		%%%%%%############(############((((((((####((######(,,,,,,*#(((((((((((((((((((((##((#####%#%#%%%%%%
		%####%#########(##%%%########(((((((((((((((((((###(****,/##%%%#####(((((((((((((((((#((####%%%#%%%%
		############%#((((#%%%#%%#####((((((((((((((((#%%%&(****/#%%%%####(((/(/((((((((((((((((((######%%%%
		############%##((%%%%##########((((((((((########%/,*,,/%%%%%%####(((((/(((((((((((((((((((#######%%
		##########%%%###(##%%(((((#####(((((((((##%%%%#%#....,,(%%%%%%####%%%#((((((((((((((##((((((##%((#%(
		#####%#%##%%%%#%###%%#(#((######(((((((((####%%%#.......(###%%&&&%#(((((((((((###(((#(((((/((#(#(((#
		%%%#######%%%%%####(####((########((((((((((##%###/..*(#(((#####((((((((((((((###((#((((/(((#####((#
		#%%##%####%%%%%%%###%####(##########((((((((#########(((((##(((((((((((((((((##(((#((((((((#%#/(((((
		##(#%#######%%%%%%%%#############################((((%#(/(###(((((((((#((#((((((((#(((((((#(((((((##
		((((((#%%####%%%%%%%%############################(/(##%%#####(((((#######((#((((((((((((((#((((##%##
		%########%###%%%%%%%###########################%%%&%%%%###((#((((########((((((((((((((###%##%%#####
		#%%%#%#####%#%%#*#%%%#%%###########%##############%&%%%%%%#(((((##########((((#(((((####%%&&%#######
		(((%%#######%%%(//(%%%%%%%%%#%%%%%%%%%%%#########%%%%####%##((((((############(#(###########(((((###
		(((((########%%((/(#%%%%%%%%%%%%%%%%%#%%%%######%%%#(#####((((((((((########################(((((###
		(((((((#%#####(##//((#%&%%%%%%%%%%%%%%%%%%####%%%%%%%#%%######(((((((##################%#####(######
		(##(((((((((##((###((##%&%%%%%%%%%%%%%%%%%%%%%%%%%%%%&&%%#####((((((#############%%%%%%%%###########
		(((##((##(((#########((#%&&&%%%%%%%%%%%%%%%%&%%%&%&%&&&%%#######((#(((#####%%%%%%%%%%%%%#######(####
		(((((((#####(#((##(((####&&&&&&%%%%%%%%%%%%%%# . **..*&%%%#############%%%%%%%%&&%%%%%%%############
		#((((((((#######(((((((###%&&&&&&%&%%%%%%%%%&..... ../&%%%%%%%###%%%%%%%&&&&&&&&%%%%%%%%############
		##########(((##((#(((((##%,,,,,*#%&&&&&&&&%&,,.....  (&&&&&&%%%%%%%&&&&&&&&&&&&%%%%%%%%#############
		############(###########%%,,,,,,,,*,,((((/,,,,,,......(&&&&&&&&&&&&&&&&&&&&%%%%%%%%%%%%#############
		########################%%........,,,,***,**,***,,,,,.../#/(#%%%%##(*,,,*%%%%%%%%%%%%%##############
		%%#############((########(  .........,,,,**********,,,,,,.......    .....,%%%%%%%%%#%%############/%
		####%#%#################%*...............,,,,,**********,,,,,............../%%%%%%%%%#%########(#%#/
		%#####%##%##############%,*,,,,......  ........,,,********,,,*,,,........... #%%%%%%%%%%%###(/%%#///
		%%######################%,*****,,,,...... ..............,,,,,*****,,,,,...... /%%%%%%%%%%#/#%%//////
		%&%%%######(##############,*************,,,,,...............,,,****,,***,,,,,./%%%%%%%#/(%%#/(////((
		%%%%&%%%###((##############,,,,,,,,***********,,,,,,....,......,,,,,,,*******#%#(#%%((%%#//(////////
		%%%%&&%%%%%##(((########(##%,......,,,,,************,,,,,,,........,,,,,,***#%%((/(%%#(((%/(/////((/
	`
	fmt.Println(kissing)
	color.HiCyan("\n\n						Nike Order Checker | Developed By jc2#9899")

	wg := sync.WaitGroup{}

	tasks, err := ReadCsv()
	if err != nil {
		color.Red("[ERR] Fatal Error Reading Tasks :%s, Closing Program.\n", err)
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}

	Import()

	for _, task := range tasks {
		wg.Add(1)
		go exeCute(task, &wg)
	}
	wg.Wait()
	color.Green("[SUCCESS] All Operations Complete. Closing in 3s...")
	color.HiCyan("[STATS] Total Success: %d\n", successInt)
	color.Red("[STATS] Total Errors: %d\n", errInt)
	time.Sleep(time.Second * 3)
}

func exeCute(Task Task, wg *sync.WaitGroup) {

	taskProxy, err := Get()
	if err != nil {
		color.Red("[ERR] Fatal Error Grabbing Proxies: %s, Closing Program.\n", err)
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}

	for {
		Task.InitClient(taskProxy)
		body, err := Task.grabOrder()
		if err != nil {
			color.Red("[ERROR] exeCute.grabOrders ERROR GRABBING ORDER.\n")
			errInt++
			handleError()
			break
		}

		if len(body.Group) > 0 {
			writeExport(body)
		}

		if err == nil {
			successInt++
			break
		}
	}
	wg.Done()
}
