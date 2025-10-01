package performance

import (
	"chezmoi-tui/internal/chezmoi"
	"chezmoi-tui/internal/integration"
	"testing"
	"time"
)

// BenchmarkIntegrationGetStatus benchmarks the GetStatus function
func BenchmarkIntegrationGetStatus(b *testing.B) {
	integ, err := integration.New()
	if err != nil {
		b.Skipf("Skipping benchmark: %v", err) // Skip if chezmoi isn't available
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := integ.GetStatus()
		if err != nil {
			// Don't fail the benchmark if chezmoi isn't initialized
			b.Logf("GetStatus returned error: %v", err)
		}
	}
}

// BenchmarkChezmoiParseStatusOutput benchmarks the status parsing function
func BenchmarkChezmoiParseStatusOutput(b *testing.B) {
	// Sample status output
	sampleOutput := ""
	for i := 0; i < 100; i++ {
		sampleOutput += " M .file" + string(rune('0'+i)) + "\\n"
	}

	client := &chezmoi.Chezmoi{} // Note: This will need to be imported properly

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.ParseStatusOutput(sampleOutput)
	}
}

// TestResponseTime tests that operations complete within acceptable time limits
func TestResponseTime(t *testing.T) {
	integ, err := integration.New()
	if err != nil {
		t.Skipf("Skipping test: %v", err) // Skip if chezmoi isn't available
	}

	t.Run("GetConfigDataResponseTime", func(t *testing.T) {
		start := time.Now()
		_, err := integ.GetConfigData()
		duration := time.Since(start)

		if err != nil {
			t.Logf("GetConfigData returned error: %v", err)
		}

		// We expect the operation to complete in under 5 seconds
		if duration > 5*time.Second {
			t.Errorf("GetConfigData took too long: %v", duration)
		} else {
			t.Logf("GetConfigData completed in: %v", duration)
		}
	})

	t.Run("GetManagedFilesResponseTime", func(t *testing.T) {
		start := time.Now()
		_, err := integ.GetManagedFiles()
		duration := time.Since(start)

		if err != nil {
			t.Logf("GetManagedFiles returned error: %v", err)
		}

		// We expect the operation to complete in under 5 seconds
		if duration > 5*time.Second {
			t.Errorf("GetManagedFiles took too long: %v", duration)
		} else {
			t.Logf("GetManagedFiles completed in: %v", duration)
		}
	})
}
