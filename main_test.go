package main

import "testing"

func BenchmarkCreateKeyByConcatenation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateKeyByConcatenation("bucket", "itemType", "itemID")
	}
}
func BenchmarkCreateKeyByBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateKeyByBytesBuffer("bucket", "itemType", "itemID")
	}
}
func BenchmarkCreateKeyByBytesWithousLen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateKeyByBytesBufferWithoutLen("bucket", "itemType", "itemID")
	}
}
