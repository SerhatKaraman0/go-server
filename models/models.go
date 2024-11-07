package models

import "time"


type User struct {
    UserID       int       `json:"user_id" db:"user_id"`
    Username     string    `json:"username" db:"username"`
    PasswordHash string    `json:"password_hash" db:"password_hash"`
    Role         string    `json:"role" db:"role"`
    DateCreated  time.Time `json:"date_created" db:"date_created"`
}

type Student struct {
    StudentID        int         `json:"student_id" db:"student_id"`
    FirstName        string      `json:"first_name" db:"first_name"`
    LastName         string      `json:"last_name" db:"last_name"`
    DOB              time.Time   `json:"dob" db:"dob"`
    EnrollmentDate   time.Time   `json:"enrollment_date" db:"enrollment_date"`
    ContactInfo      interface{} `json:"contact_info" db:"contact_info"` 
    EncryptedRecords []byte      `json:"encrypted_records" db:"encrypted_records"`
}

type Grade struct {
    GradeID     int       `json:"grade_id" db:"grade_id"`
    StudentID   int       `json:"student_id" db:"student_id"`
    CourseCode  string    `json:"course_code" db:"course_code"`
    GPA         float64   `json:"gpa" db:"gpa"`
    CGPA        float64   `json:"cgpa" db:"cgpa"`
    LastUpdated time.Time `json:"last_updated" db:"last_updated"`
}

type Document struct {
    DocumentID   int       `json:"document_id" db:"document_id"`
    StudentID    int       `json:"student_id" db:"student_id"`
    DocumentType string    `json:"document_type" db:"document_type"`
    DocumentData []byte    `json:"document_data" db:"document_data"`
    RsaSignature []byte    `json:"rsa_signature" db:"rsa_signature"`
    DateCreated  time.Time `json:"date_created" db:"date_created"`
    Status       string    `json:"status" db:"status"`
}

type Invoice struct {
    InvoiceID  int       `json:"invoice_id" db:"invoice_id"`
    StudentID  int       `json:"student_id" db:"student_id"`
    DocumentID int       `json:"document_id" db:"document_id"`
    AmountDue  float64   `json:"amount_due" db:"amount_due"`
    Status     string    `json:"status" db:"status"`
    DateIssued time.Time `json:"date_issued" db:"date_issued"`
}

type Payment struct {
    PaymentID    int       `json:"payment_id" db:"payment_id"`
    InvoiceID    int       `json:"invoice_id" db:"invoice_id"`
    AmountPaid   float64   `json:"amount_paid" db:"amount_paid"`
    PaymentDate  time.Time `json:"payment_date" db:"payment_date"`
    ReceiptData  []byte    `json:"receipt_data" db:"receipt_data"`
    MacCode      []byte    `json:"mac_code" db:"mac_code"`
}

type Key struct {
    KeyID        int       `json:"key_id" db:"key_id"`
    KeyType      string    `json:"key_type" db:"key_type"`
    KeyValue     []byte    `json:"key_value" db:"key_value"`
    CreatedBy    int       `json:"created_by" db:"created_by"`
    LastUpdated  time.Time `json:"last_updated" db:"last_updated"`
}

type ActivityLog struct {
    LogID     int       `json:"log_id" db:"log_id"`
    UserID    int       `json:"user_id" db:"user_id"`
    Action    string    `json:"action" db:"action"`
    Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

type Request struct {
    RequestID    int       `json:"request_id" db:"request_id"`
    StudentID    int       `json:"student_id" db:"student_id"`
    DocumentType string    `json:"document_type" db:"document_type"`
    RequestDate  time.Time `json:"request_date" db:"request_date"`
    Status       string    `json:"status" db:"status"`
}

