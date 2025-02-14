# Image Formater Program

This Go program allows you to resize images based on predefined formats like license, driver, and national ID, and supports three different DPI values (300, 600, and 1200) for high-quality prints. It provides an easy way to resize images for printing, ensuring they meet specific format requirements.

## Features

- Resize images to predefined formats: `license`, `driver`, `national_id`
- Supports three DPI options for resizing: `300`, `600`, `1200`
- Displays available format sizes before running the program
- Supports `jpeg` and `png` formats for input and output

## Installation

1. Ensure that you have **Go** installed on your machine. If not, you can download it from the official website: [Go Download](https://golang.org/dl/).

2. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/image-resizer.git
   ```
   
   ```bash
   cd image-resizer
   ```

3. Install the required Go modules (if any):

   ```bash
   go mod tidy
   ```

## Usage

You can either view available format details or resize an image with the following commands.

### View Available Format Details

To see the available formats and their sizes before processing an image, run the following command:

```bash
go run main.go -formats
```

This will output the available formats and their dimensions, for example:

```
Available formats and their sizes:
license: 3.50 cm x 4.50 cm
driver: 5.00 cm x 5.00 cm
national_id: 2.50 cm x 3.00 cm
```

### Resize an Image

To resize an image, you need to specify the input image file, the desired format, and the DPI value. Here's how you can run the program:

```bash
go run main.go -img <image_path> -format <license|driver|national_id> -dpi <300|600|1200>
```

#### Example:

```bash
go run main.go -img input.jpg -format license -dpi 600
```

This will resize the `input.jpg` image to the dimensions of a license with a DPI of 600 and save the output in the same directory.

### Flags:

- `-img <image_path>`: Specifies the path to the input image file.
- `-format <license|driver|national_id>`: Specifies the format for resizing.
- `-dpi <300|600|1200>`: Specifies the DPI (resolution) for resizing. Choose from 300, 600, or 1200 DPI.
- `-formats`: Displays available format sizes (without running the resizing task).

## Example Output

After resizing, the program will print the output path of the resized image:

```
Resized image saved: input_license_600.jpg
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


### Explanation:

- **What it does**: The program resizes images based on predefined formats and DPI values.
- **How to run**: Step-by-step guide on how to run the program with flags.
- **Flags explained**: Each flag's purpose is explained to help users understand how to use them.
- **Example**: Shows an example of how to run the program with both the `-formats` and the resizing flags.
