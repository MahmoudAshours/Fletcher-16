# **Fletcher Checksum in Go**

## **Overview**
This Go program reads a file (`GameTheory.pdf`), calculates its **Fletcher-16 checksum**, and prints the result.

The **Fletcher checksum** is an error-detecting code that improves upon simple checksums by reducing the probability of undetected errors.

---

## **How It Works**
The program consists of three main functions:
- **`main()`**: Calls `readFile()` to read the file and then computes the checksum using `fletcher16()`.
- **`fletcher16(data []uint8, count int) uint16`**: Implements the Fletcher-16 checksum algorithm.
- **`readFile() []byte`**: Reads the file `GameTheory.pdf` and returns its contents as a byte array.

---

## **Code Explanation**

### **1. `main()` Function**
```go
func main() {
	byteArr := readFile()
	fmt.Printf("Fletcher checksum , %04x\n", fletcher16(byteArr, len(byteArr)))
}
```
- Calls `readFile()` to read the contents of `GameTheory.pdf`.
- Passes the byte array and its length to `fletcher16()`.
- Prints the computed Fletcher-16 checksum in **hexadecimal format**.

---

### **2. `fletcher16()` Function**
```go
func fletcher16(data []uint8, count int) uint16 {

	var a uint16 = 0
	var b uint16 = 0

	for i := 0; i < count; i++ {
		a = (a + uint16(data[i])) % 255
		b = (b + a) % 255
	}
	return (b << 8) | a
}
```
#### **How Fletcher-16 Works:**
1. Initializes two 8-bit sums: `a` and `b`.
2. Iterates through the data:
   - `a = (a + data[i]) % 255` → Adds each byte to `a`, keeping it within 8 bits.
   - `b = (b + a) % 255` → Accumulates `a` values, ensuring 8-bit bounds.
3. Combines `b` (high byte) and `a` (low byte) into a **16-bit checksum** (`(b << 8) | a`).

#### **Why Modulo 255?**
- Helps detect burst errors and prevents overflow issues.
- Ensures values remain within 8-bit constraints for portability.

---

### **3. `readFile()` Function**
```go
func readFile() []byte {
	data, err := os.ReadFile("GameTheory.pdf")
	if err != nil {
		return []byte{}
	}
	return data
}
```
#### **How It Works:**
- Uses `os.ReadFile()` to **read the entire file into memory**.
- If the file cannot be read (e.g., missing file), it returns an empty byte array (`[]byte{}`).
- Otherwise, it returns the file contents.

#### **Limitations:**
- This method is not memory-efficient for **very large files**.
- A better approach for large files is reading in chunks using `bufio.Reader`.

---

## **Example Output**
If `GameTheory.pdf` exists and is successfully read, the output will be:
```sh
Fletcher checksum , 1bea
```
(Checksum value will vary based on file content.)

If the file is missing, the program will silently return `0` as the checksum.

---

## **Possible Enhancements**
1. **Better Error Handling:**
   - Log errors instead of returning an empty array in `readFile()`.
   - Example:
     ```go
     if err != nil {
         fmt.Println("Error reading file:", err)
         return nil
     }
     ```

2. **Processing Large Files Efficiently:**
   - Read in chunks using `bufio.Reader` instead of `os.ReadFile()`.
   - This prevents high memory usage for large files.

3. **Making the File Name Configurable:**
   - Allow users to specify the file via command-line arguments.
   - Example:
     ```go
     func main() {
         if len(os.Args) < 2 {
             fmt.Println("Usage: go run main.go <filename>")
             return
         }
         filename := os.Args[1]
         byteArr := readFile(filename)
         fmt.Printf("Fletcher checksum , %04x\n", fletcher16(byteArr, len(byteArr)))
     }
     ```

---

## **Conclusion**
This program demonstrates the **Fletcher-16 checksum** for file integrity verification. It efficiently calculates a checksum but can be improved with better error handling and memory-efficient file reading. 🚀

---

## **References**
- [Fletcher Checksum - Wikipedia](https://en.wikipedia.org/wiki/Fletcher%27s_checksum)
- [Go os.ReadFile Documentation](https://pkg.go.dev/os#ReadFile)

