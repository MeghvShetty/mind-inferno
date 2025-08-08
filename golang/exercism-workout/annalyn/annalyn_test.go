ls
package annalyn

import "testing"

func TestCanFastAttack(t *testing.T) {
	tests := []struct {
		name           string
		knightIsAwake  bool
		expectedResult bool
	}{
		{
			name:           "knight is sleeping",
			knightIsAwake:  false,
			expectedResult: true,
		},
		{
			name:           "knight is awake",
			knightIsAwake:  true,
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanFastAttack(tt.knightIsAwake)
			if result != tt.expectedResult {
				t.Errorf("CanFastAttack(%v) = %v; want %v", tt.knightIsAwake, result, tt.expectedResult)
			}
		})
	}
}

func TestCanSpy(t *testing.T) {
	tests := []struct {
		name            string
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		expectedResult  bool
	}{
		{
			name:            "all characters are sleeping",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			expectedResult:  false,
		},
		{
			name:            "knight is awake",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			expectedResult:  true,
		},
		{
			name:            "archer is awake",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			expectedResult:  true,
		},
		{
			name:            "prisoner is awake",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			expectedResult:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanSpy(tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake)
			if result != tt.expectedResult {
				t.Errorf("CanSpy(%v, %v, %v) = %v; want %v",
					tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, result, tt.expectedResult)
			}
		})
	}
}

func TestCanSignalPrisoner(t *testing.T) {
	tests := []struct {
		name            string
		archerIsAwake   bool
		prisonerIsAwake bool
		expectedResult  bool
	}{
		{
			name:            "archer is awake, prisoner is awake",
			archerIsAwake:   true,
			prisonerIsAwake: true,
			expectedResult:  false,
		},
		{
			name:            "archer is sleeping, prisoner is awake",
			archerIsAwake:   false,
			prisonerIsAwake: true,
			expectedResult:  true,
		},
		{
			name:            "archer is sleeping, prisoner is sleeping",
			archerIsAwake:   false,
			prisonerIsAwake: false,
			expectedResult:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanSignalPrisoner(tt.archerIsAwake, tt.prisonerIsAwake)
			if result != tt.expectedResult {
				t.Errorf("CanSignalPrisoner(%v, %v) = %v; want %v",
					tt.archerIsAwake, tt.prisonerIsAwake, result, tt.expectedResult)
			}
		})
	}
}

func TestCanFreePrisoner(t *testing.T) {
	tests := []struct {
		name            string
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		petDogIsPresent bool
		expectedResult  bool
	}{
		{
			name:            "prisoner is awake, others are sleeping, no dog",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			petDogIsPresent: false,
			expectedResult:  true,
		},
		{
			name:            "prisoner is awake, knight is awake, archer is sleeping, no dog",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			petDogIsPresent: false,
			expectedResult:  false,
		},
		{
			name:            "dog is present, archer is sleeping",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			petDogIsPresent: true,
			expectedResult:  true,
		},
		{
			name:            "dog is present, archer is awake",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			petDogIsPresent: true,
			expectedResult:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanFreePrisoner(tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, tt.petDogIsPresent)
			if result != tt.expectedResult {
				t.Errorf("CanFreePrisoner(%v, %v, %v, %v) = %v; want %v",
					tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, tt.petDogIsPresent, result, tt.expectedResult)
			}
		})
	}
}
