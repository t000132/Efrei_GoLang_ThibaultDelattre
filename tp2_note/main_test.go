package main

import (
	"strings"
	"testing"
	"time"
)

// Test EmailNotifier
func TestEmailNotifier(t *testing.T) {
	email := EmailNotifier{Address: "test@example.com"}
	
	err := email.Send("Test message")
	if err != nil {
		t.Errorf("EmailNotifier.Send() failed: %v", err)
	}
	
	if email.GetType() != "Email" {
		t.Errorf("Expected type 'Email' => got '%s'", email.GetType())
	}
}

// Test de SMSNotifier avec numéro valide
func TestSMSNotifier_Valid(t *testing.T) {
	sms := SMSNotifier{PhoneNumber: "0610203040"}
	
	err := sms.Send("Test SMS")
	if err != nil {
		t.Errorf("SMSNotifier.Send() failed with valid number: %v", err)
	}
}

// Test SMSNotifier avec numéro invalide
func TestSMSNotifier_Invalid(t *testing.T) {
	sms := SMSNotifier{PhoneNumber: "0712345678"}
	
	err := sms.Send("Test SMS")
	if err == nil {
		t.Error("SMSNotifier.Send() should fail with invalid number")
	}
	
	if !strings.Contains(err.Error(), "invalide") {
		t.Errorf("Error message should contain 'invalide' => got: %v", err)
	}
}

// Test PushNotifier
func TestPushNotifier(t *testing.T) {
	push := PushNotifier{DeviceID: "device123"}
	
	err := push.Send("Test push")
	if err != nil {
		t.Errorf("PushNotifier.Send() failed: %v", err)
	}
	
	if push.GetType() != "Push" {
		t.Errorf("Expected type 'Push' => got '%s'", push.GetType())
	}
}

// Test MemoryStorer
func TestMemoryStorer(t *testing.T) {
	storer := &MemoryStorer{}
	
	// Test Store
	record := NotificationRecord{
		Type:      "Email",
		Message:   "Test",
		Timestamp: time.Now(),
	}
	
	err := storer.Store(record)
	if err != nil {
		t.Errorf("MemoryStorer.Store() failed: %v", err)
	}
	
	// Test GetAll
	records := storer.GetAll()
	if len(records) != 1 {
		t.Errorf("Expected 1 record => got %d", len(records))
	}
	
	if records[0].Type != "Email" {
		t.Errorf("Expected type 'Email' => got '%s'", records[0].Type)
	}
}