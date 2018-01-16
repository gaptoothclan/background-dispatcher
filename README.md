#Background Dispatcher

Create a dispatcher with the amount of workers you need and the max size of your queue.

Run the dispatcher with Run()

Below I am adding a function to the dispatcher to be completed in the background

```
	dispatcher := NewDispatcher(5, 20)
	dispatcher.Run()

  for x := 0; x < 20; x++ {
		dispatcher.AddToQueue(func() {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("Func DONE")
		})
	}
```