
### Common Usages

+ **Managing Goroutines**


زمانی که یک حلقه ی بی نهایت قرار است به صورت گریس فول خاتمه یابد یا فرمان دللاین یا کنسل بیاد

```go


func longRunningTask(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Task canceled")
            return
        default:
            // Do work
        }
    }
}

```
+ **Timeouts**

در جاهایی از برنامه که احتمال تاخیر در جواب باشد مانند database , network در این موارد پاسخ ممکن است با تاخیر باشد پس می توان تایمر گذاشت :

```go
func timeoutExample() {
    // Create a context with a timeout of 2 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Ensure the context is canceled to avoid resource leaks

    // Start a goroutine that does some work
    done := make(chan struct{})
    go func() {
        defer close(done)
        select {
        case <-time.After(3 * time.Second):
            fmt.Println("Work completed")
        case <-ctx.Done():
            fmt.Println("Work timed out:", ctx.Err())
        }
    }()

    // Wait for goroutine to exit
    <-done
}
```


+ **graceful shutdown**

زمانی که برنامه به صورت طبیعی بخواد شات داون شه بهتره به تمامی فانکشن هایی که در حال اجرا هستن بگیم که با کمترین خسارت بسته بشن

```go
func main() {
    // Create a cancellable context
    ctx, cancel := context.WithCancel(context.Background())

    // Set up a channel to listen for OS signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    // مثلن کدی که در مثال اول آورده شده
    go longRunningTask(ctx)

    // Block until a signal is received
    sig := <-sigChan
    fmt.Printf("Received signal: %s\n", sig)

    // Cancel the context to notify goroutines to stop
    cancel()

    // Give goroutines time to clean up before exiting
    time.Sleep(5 * time.Second)
    fmt.Println("Shutting down gracefully")
}
```

+ **context.WithValue**

   میشه به جای استفاده از متغییر سراری - گلوبال وریبل - از این استاده کرد ، و از اونجا به پایین نمی تونن روی گلوبال تاثیر بزارن


ولی در کل بدیش اینه که **Side Effects** داریم یعنی داده از جای نا معلوم در دسترسه و میشه گفت **dependency injection** تولید می کنه


+ **cancel()**

توجه شود حتما بعد از همه چی کلوز بشه وگرنه مموری لیک رخ میده
