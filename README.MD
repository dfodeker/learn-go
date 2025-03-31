# Go Programming Language Exercises

This repository contains exercises to practice the Go programming language, primarily based on:

- **The Go Programming Language** by Alan A. A. Donovan & Brian W. Kernighan
- **Boot.dev's Go Course** – [Check it out here](https://www.boot.dev/courses/learn-golang), highly recommended for anyone learning Go!

## 📚 Contents

### Chapter 1: Getting Started
1. **1.1** – Modify the `echo` program to also print `os.Args[0]`, the name of the command that invoked it.
2. **1.2** – Modify the `echo` program to print the index and value of each of its arguments, one per line.
3. **1.3** – Experiment to measure the difference in running time between inefficient versions and the one using `strings.Join`. (See Section 1.6 for `time` package and Section 11.4 for benchmark tests.)
4. **1.4** – Modify `dup2` to print the names of all files where each duplicated line occurs.

### Boot.dev Exercises
Exercises from the **Boot.dev Go Course**, covering fundamental and advanced Go topics:

- **ch_4_structs** – Working with structs
- **ch_5_interfaces** – Understanding interfaces in Go
- **ch_8_slices** – Manipulating slices efficiently
- **ch_9_maps** – Working with maps
- **ch_10_packages** – Understanding and using packages
- **ch_11_pointers** – Pointer concepts in Go
- **ch_12_channels** – Concurrency with channels

## 🚀 Getting Started

To run the exercises, make sure you have Go installed on your system.

1. Clone this repository:
   ```sh
   git clone https://github.com/dfodeker/learn-go.git
   cd learn-go

   ```
2. Navigate to the exercise directory you want to run.
3. Run the Go program:
   ```sh
   go run main.go
   ```
4. Follow the instructions in the code comments to attempt to re-solve the problem.
5. For testing, you can run:
   ```sh
   go test
   ```
## 📌 Notes
The exercises are based on personal learning and practice.

Contributions are welcome! Feel free to open issues or pull requests.

## 🏆 Acknowledgments
Special thanks to:

Alan A. A. Donovan & Brian W. Kernighan for their fantastic book.

Boot.dev for their well-structured Go course.
ps.Sorry I didn't pay for the course, but I still want to learn Go.
