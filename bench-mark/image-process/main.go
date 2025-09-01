package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const FolderPath = `E:/Downloads/plantvillage dataset/Pepper,_bell___Bacterial_spot`
const Threshold = 128

func main() {
	files, err := os.ReadDir(FolderPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d files\n", len(files))

	// --- Measure ThresholdIf ---
	fmt.Println("Before ThresholdIf:")
	PrintMemUsage()

	start := time.Now()
	var totalBytesIf uint64
	var lastGrayIf *image.Gray
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		path := filepath.Join(FolderPath, f.Name())
		img, err := LoadImage(path)
		if err != nil {
			fmt.Println("LoadPNG error:", path, err)
			continue
		}
		gray := ThresholdIf(img, Threshold)
		lastGrayIf = gray
		totalBytesIf += uint64(len(gray.Pix))
	}
	elapsedIf := time.Since(start)
	fmt.Printf("ThresholdIf elapsed: %v, approx memory: %d KB\n", elapsedIf, totalBytesIf/1024)
	PrintMemUsage()

	// --- Measure ThresholdBranchless ---
	fmt.Println("Before ThresholdBranchless:")
	PrintMemUsage()

	start = time.Now()
	var totalBytesBranchless uint64
	var lastGrayBranchless *image.Gray
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		path := filepath.Join(FolderPath, f.Name())
		img, err := LoadImage(path)
		if err != nil {
			fmt.Println("LoadPNG error:", path, err)
			continue
		}
		gray := ThresholdBranchless(img, Threshold)
		lastGrayBranchless = gray
		totalBytesBranchless += uint64(len(gray.Pix))
	}
	elapsedBranchless := time.Since(start)
	fmt.Printf("ThresholdBranchless elapsed: %v, approx memory: %d KB\n", elapsedBranchless, totalBytesBranchless/1024)
	PrintMemUsage()

	// Save last results
	if lastGrayIf != nil {
		_ = SaveGrayImage(lastGrayIf, "./last_if.png")
		fmt.Println("Saved last ThresholdIf result to ./last_if.png")
	}
	if lastGrayBranchless != nil {
		_ = SaveGrayImage(lastGrayBranchless, "./last_branchless.png")
		fmt.Println("Saved last ThresholdBranchless result to ./last_branchless.png")
	}
}

// ---------------------
// Helper functions
// ---------------------

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func SaveGrayImage(img *image.Gray, path string) error {
	if img == nil {
		return nil
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

// Hàm in thông tin bộ nhớ
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KB\tTotalAlloc = %v KB\tSys = %v KB\tNumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

// ThresholdIf
func ThresholdIf(img image.Image, threshold uint8) *image.Gray {
	bounds := img.Bounds()
	result := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((r + g + b) / 3 >> 8)
			if gray > threshold {
				result.SetGray(x, y, color.Gray{255})
			} else {
				result.SetGray(x, y, color.Gray{0})
			}
		}
	}
	return result
}

// ThresholdBranchless
func ThresholdBranchless(img image.Image, threshold uint8) *image.Gray {
	bounds := img.Bounds()
	result := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((r + g + b) / 3 >> 8)
			diff := int(gray) - int(threshold)
			mask := uint8((diff >> 31) ^ 1) // 1 nếu gray>threshold, 0 nếu <=
			result.SetGray(x, y, color.Gray{mask * 255})
		}
	}
	return result
}
