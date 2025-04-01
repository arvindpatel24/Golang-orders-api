package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/arvindpatel24/orders-api/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel() // here defer means cancel will be called after the end of this function(main())

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
