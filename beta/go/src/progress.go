// progress.go
package main

import (
	"fmt"
	"time"
)

type BuildProgress struct {
	currentStep string
	startTime   time.Time
}

func NewBuildProgress() *BuildProgress {
	return &BuildProgress{
		startTime: time.Now(),
	}
}

func (bp *BuildProgress) StartBuildProcess() {
	bp.currentStep = "Initializing build"
	fmt.Println("🚀 Build process started...")
}

func (bp *BuildProgress) UpdateProgress(step string) {
	bp.currentStep = step
	fmt.Printf("⏳ Progress: %s...\n", step)
}

func (bp *BuildProgress) CompleteBuildProcess(binaryPath string) {
	duration := time.Since(bp.startTime)
	fmt.Println("✅ Build process completed successfully!")
	fmt.Printf("📦 Binary created at: %s\n", binaryPath)
	fmt.Printf("⏱️  Total build time: %v\n", duration)
}

func (bp *BuildProgress) HandleBuildError(err error) {
	fmt.Printf("❌ Build process failed: %v\n", err)
	fmt.Printf("Current step: %s\n", bp.currentStep)
}
