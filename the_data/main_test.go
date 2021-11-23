package main

import "testing"

func TestPalindromeP(t *testing.T) {
	if !palindromeP("平a man, a plan, a canal, Panama.平") {
		t.Error("this should have detected as a palindrome")
	}

}

func TestPalindromeP_notPalindrome(t *testing.T) {
	if palindromeP("Rudderford") {
		t.Error("this should not have detected as a palindrome")
	}

}
