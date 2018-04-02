package users

import "testing"

func TestIsValidUsername(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"sim", true},
		{"simulation123", false},
		{"lo", false},
		{"l", false},
		{"bulldog", true},
		{"dog", false},
		{"DOG", false},
		{"catfish", true},
		{"CATfish", true},
		{"catFISH", true},
		{"scatter", true},
		{"sCATter", true},
		{"SCATTER", true},
		{"cat", false},
		{"CAT", false},
		{"seahorse", true},
		{"seaHORSE", true},
		{"SEAhorse", true},
		{"horse", false},
		{"doogCAT", false},
		{"DOGcatfish", false},
		{"doogcatfish", true},
	}

	for _, test := range tests {
		var u User
		u.Username = test.input

		if got := u.IsValidUsername(); got != test.want {
			t.Errorf("u.Username = %q u.IsValidUsername() = %v", test.input, got)
		}
	}
}

func TestIsValidPhoneNumber(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"09898", false},
		{"0ar9898", false},
		{"08992221110", true},
		{"98903278114", true},
		{"08992221110989889", false},
	}

	for _, test := range tests {
		var u User
		u.PhoneNumber = test.input

		if got := u.IsValidPhoneNumber(); got != test.want {
			t.Errorf("u.PhoneNumber = %q u.IsValidPhoneNumber() = %v", test.input, got)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"x@test.com", true},
		{"simon.rouger@test.com", true},
		{"sim..@test.com", false},
		{"sim@@test.com", false},
		{"sim@test@com", false},
		{"sim@test..com", false},
	}

	for _, test := range tests {
		var u User
		u.Email = test.input

		if got := u.IsValidEmail(); got != test.want {
			t.Errorf("u.Email = %q u.IsValidEmail() = %v", test.input, got)
		}
	}
}
