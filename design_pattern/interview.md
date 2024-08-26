### انسانی

+ حتما درباره ی soft skill صحبت شه ، بگیم که شرکت و دوست داریم ، بگیم نمیزارم جو سمی شه ، میفهمم trade of کد کامپلکس و راحتی رو

+ تعریف از لیدر و اطاعت

+ این که هر وقت همه می خوان برن یا مهاجرت یا از شرکت ، من پایبند  شرکتم و احترام میزارم به سرمایه گذار




+ website
Pramp or Interviewing.io

### time complexity

+ **O(n)**

در کد زیر چون حلقه یک بار تا ته باید طی شه  ، به اندازه طول حلقه است

```go
package main

import "fmt"

// Function to find the maximum value in a slice
func findMax(nums []int) int {
    if len(nums) == 0 {
        return -1 // or handle empty slice case
    }

    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}

func main() {
    nums := []int{3, 5, 7, 2, 8, 9, 1}
    fmt.Println("Max:", findMax(nums))
}
```

+ **Bubble Sort O(n²)**

```go

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
```

+ **Merge Sort O(n log n)**

```go
func mergeSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    return merge(left, right)
}

func merge(left, right []int) []int {
    result := []int{}
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    return result
}


```

+ **Binary Search O(log n)**

```go
func binarySearch(arr []int, target int) int {
    low, high := 0, len(arr)-1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1 // Target not found
}


```