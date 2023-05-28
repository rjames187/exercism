// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	res := Resident{name, age, address}
	return &res
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	noName := r.Name == ""
	noAddress := r.Address["street"] == ""
	if noName || noAddress {
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, res := range residents {
		if res.HasRequiredInfo() {
			count++
		}
	}
	return count
}
