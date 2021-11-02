// Golang program to illustrate the usage of 
// AfterFunc() function 
  
// Including main package 
package main 
  
// Importing fmt and time 
import ( 
    "fmt"
    "time"
) 
  
// Main function 
func main() { 
  
    // Defining duration parameter of 
    // AfterFunc() method 
    DurationOfTime:= time.Duration(3) * time.Second 
  
    // Defining function parameter of 
    // AfterFunc() method 
    f:= func() { 
  
        // Printed when its called by the 
        // AfterFunc() method in the time 
        // stated above 
        fmt.Println("Function called by "+ 
            "AfterFunc() after 3 seconds") 
    } 
  
    // Calling AfterFunc() method with its 
    // parameter 
    Timer1:= time.AfterFunc(DurationOfTime, f) 
  
    // Calling stop method  
    // w.r.to Timer1 
    defer Timer1.Stop() 
  
    // Calling sleep method 
    time.Sleep(10 * time.Second) 
}
