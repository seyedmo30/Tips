اگر بخوهیم از قابلیت cancel , Done کانتکست استفاده کنیم میبایست تمامی فانکشن هایی که این پارامتر را میگیرند ، استفاده کنیم :

```


func F(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("F: received cancellation signal")
		return
	default:
		// Simulate some work
		time.Sleep(1 * time.Second)
		fmt.Println("F: working...")
	}
}


```

این یعنی همیشه ابتدا چک کنیم کانتکس دان شده یا خیر ، همچنین باید هر جا که رانتایم منمتظر می ماند ، دان شدن را چک کنیم

