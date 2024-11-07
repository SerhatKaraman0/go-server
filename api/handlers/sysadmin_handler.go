package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SerhatKaraman0/go-server/models"
	"github.com/SerhatKaraman0/go-server/utils"
)

var student models.Student

func CreateStudentAccount(w http.ResponseWriter, r *http.Request) {
	
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("Parsed student data: %+v\n", student)

	if student.FirstName == "" || student.LastName == "" {
		http.Error(w, "First name and last name are required", http.StatusBadRequest)
		return
	}

    if student.EnrollmentDate.IsZero() {
		student.EnrollmentDate = time.Now()
	}


	contactInfoJSON, err := json.Marshal(student.ContactInfo)
	if err != nil {
		http.Error(w, "Invalid contact info format", http.StatusBadRequest)
		return
	}

	// Parse DateOfBirth and EnrollmentDate to ensure they're in valid format
	if student.DOB.IsZero() {
		http.Error(w, "Date of birth is required", http.StatusBadRequest)
		return
	}
	if student.EnrollmentDate.IsZero() {
		student.EnrollmentDate = time.Now() // Default to the current date if not provided
	}

	// Debug: Print all parsed fields
	fmt.Printf("First Name: %s\n", student.FirstName)
	fmt.Printf("Last Name: %s\n", student.LastName)
	fmt.Printf("Contact Info: %s\n", string(contactInfoJSON))
	fmt.Printf("Date of Birth: %s\n", student.DOB.Format("2006-01-02"))
	fmt.Printf("Enrollment Date: %s\n", student.EnrollmentDate.Format("2006-01-02"))
	fmt.Printf("Encrypted Records: %v\n", student.EncryptedRecords)

	// Update the SQL query to include DateOfBirth and EnrollmentDate
	query := `
		INSERT INTO students (
			student_id,
			first_name, 
			last_name, 
			contact_info, 
			encrypted_records,
			dob,
			enrollment_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING first_name
	`

	// Execute the query with DateOfBirth and EnrollmentDate
	err = utils.DB.QueryRow(
		query,
		student.StudentID,
		student.FirstName,
		student.LastName,
		contactInfoJSON,
		student.EncryptedRecords,
		student.DOB,
		student.EnrollmentDate,
	).Scan(&student.FirstName)

	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to create student account: %v", err), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}



func UpdateStudentAccount(w http.ResponseWriter, r *http.Request) {


	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}


	if student.StudentID == 0 {
		http.Error(w, "Student ID is required", http.StatusBadRequest)
		return
	}


	query := `
		UPDATE students 
		SET first_name = $1, last_name = $2, contact_info = $3, 
			dob = $4, enrollment_date = $5
		WHERE student_id = $6
	`


	contactInfoJSON, err := json.Marshal(student.ContactInfo)
	if err != nil {
		http.Error(w, "Invalid contact info format", http.StatusBadRequest)
		return
	}


	_, err = utils.DB.Exec(
		query,
		student.FirstName,
		student.LastName,
		contactInfoJSON,
		student.DOB,
		student.EnrollmentDate,
		student.StudentID,
	)

	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to update student account: %v", err), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func DeleteStudentAccount(w http.ResponseWriter, r *http.Request) {

}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	// Prepare the SQL query to fetch all students
	query := `
		SELECT student_id, first_name, last_name, contact_info, dob, enrollment_date 
		FROM students
	`

	// Execute the query
	rows, err := utils.DB.Query(query)
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to retrieve students: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the student records
	var students []models.Student

	// Iterate through the rows
	for rows.Next() {
		var student models.Student
		var contactInfoJSON []byte // To hold the contact info JSON

		// Scan the row into the student struct
		if err := rows.Scan(&student.StudentID, &student.FirstName, &student.LastName, &contactInfoJSON, &student.DOB, &student.EnrollmentDate); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			http.Error(w, "Error reading student data", http.StatusInternalServerError)
			return
		}

		// Unmarshal the contact info from JSON
		if err := json.Unmarshal(contactInfoJSON, &student.ContactInfo); err != nil {
			fmt.Printf("Error unmarshalling contact info: %v\n", err)
			http.Error(w, "Error processing contact info", http.StatusInternalServerError)
			return
		}

		// Append the student record to the slice
		students = append(students, student)
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating rows: %v\n", err)
		http.Error(w, "Error processing student records", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Use 200 OK for a successful retrieval
	json.NewEncoder(w).Encode(students) // Encode the slice of students to JSON
}


func UpdateDesKey(w http.ResponseWriter, r *http.Request) {

}

func UpdateRsaKey(w http.ResponseWriter, r *http.Request) {
	
}
