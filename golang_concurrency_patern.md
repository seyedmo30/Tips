

# پترن جنراتور



در این پترن ، سرویسی به نام جنراتور وجود دارد که در خود آن چنل خروجی اش ایجاد شده و کلوز شدن چنل هم در خود جنراتور اتفاق می افتد ، همچنین چنل را ریترن می کند .

همچنین هر چنل می توان استراکتی باشد که در خود ارور هم دارد ، و یکی از ورودی های جنراتور ، فلگ است و جنراتور در خود سلکت دارد ، مورد اول فلگ را بررسی می کند، در صورتی که دان شود دستور مربوطه اجرا می شود و در صورتی که نباشد روال عادی طی می شود


```

type Result struct {
	data int
	err  error
}



func handler() {
	input := []int{1, 2, 3, 4, 5, 6}

	// Explicit cancellation
	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(doneCh, input)

	for data := range inputCh {
		if data == 1 {
			return
		}
	}
}



func generator(doneCh chan struct{}, input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case <-doneCh:
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}



func consumer(inputCh chan int, resultCh chan Result) {
	defer close(resultCh)

	for data := range inputCh {
		resp, err := callDatabase(data)

		result := Result{
			data: resp,
			err:  err,
		}

		resultCh <- result
	}
}


```

