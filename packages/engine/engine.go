package engine

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

// The simulated user will pseudo-randomly pick one of the actions
// the endpoint will be called, a time will be measured for the request
// results are either: request timesout/fails or the time it took to complete the request

func chooseAction(profile UserProfile) {

}

func callEndpoint(action UserAction) {
    
}

// this function continues working until the higher context tells this function to return after
// execution time has finished
func SimulateUser(ctx context.Context, profile UserProfile) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done working")
			return
		default:
			for i, action := range profile.UserActions {
				fmt.Printf("%s, %v, %f, %s \n", action.Name, action.Endpoint, profile.ActionOdds[i], profile.UserActions[i].Data)
			}
		}

        // randomly generates a timeout, either 1 or 2 seconds
        // gonna try to generate a time between 1 and 3 seconds with decimals added
		randVal := rand.IntN(3)
		if randVal <= 0 {
			randVal = 1
		}
		fmt.Printf("%d", randVal)
		time.Sleep(time.Second * time.Duration(randVal))
	}

}
