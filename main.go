package main

import (
	"log"
	"net/http"

	"github.com/SerhatKaraman0/go-server/api/handlers"
	"github.com/SerhatKaraman0/go-server/utils"
	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter().StrictSlash(true)
    
    // System Admin routes
    r.HandleFunc("/api/sysadmin/students", handlers.CreateStudentAccount).Methods("POST")
    r.HandleFunc("/api/sysadmin/students/{id}", handlers.UpdateStudentAccount).Methods("PUT")
    r.HandleFunc("/api/sysadmin/students/{id}", handlers.DeleteStudentAccount).Methods("DELETE")
    r.HandleFunc("/api/sysadmin/students", handlers.GetAllStudents).Methods("GET")
    r.HandleFunc("/api/sysadmin/keys/des", handlers.UpdateDesKey).Methods("PUT")
    r.HandleFunc("/api/sysadmin/keys/rsa", handlers.UpdateRsaKey).Methods("PUT")

    // Staff routes
    r.HandleFunc("/api/staff/{student_id}/grades", handlers.EnterGrades).Methods("POST")
    r.HandleFunc("/api/staff/invoices", handlers.SendInvoice).Methods("POST")
    r.HandleFunc("/api/staff/documents", handlers.SendTranscript).Methods("POST")
    r.HandleFunc("/api/staff/invoices", handlers.ListAllInvoices).Methods("GET")
    r.HandleFunc("/api/staff/invoices/{invoiceId}", handlers.UpdateInvoiceStatus).Methods("PUT")
    r.HandleFunc("/api/staff/students/{studentId}", handlers.ViewStudentRecords).Methods("GET")

    // Student routes
    r.HandleFunc("/api/students/{student_id}/documents/request", handlers.RequestTranscript).Methods("POST")
    r.HandleFunc("/api/students/{student_id}/invoices", handlers.GetTranscript).Methods("GET")
    r.HandleFunc("/api/students/{student_id}/grades", handlers.ViewOwnGrades).Methods("GET")
    r.HandleFunc("/api/students/{student_id}", handlers.UpdatePersonalInformation).Methods("PUT")
    r.HandleFunc("/api/students/{student_id}/payments", handlers.SubmitPaymentReceipt).Methods("POST")

    // Auth routes
    r.HandleFunc("/api/auth/login", handlers.UserLogin).Methods("POST")
    r.HandleFunc("/api/auth/logout", handlers.UserLogout).Methods("POST")

    // Security routes
    r.HandleFunc("/api/security/keys/des", handlers.GetDesKeyStatus).Methods("GET")
    r.HandleFunc("/api/security/keys/rsa", handlers.GetRsaKeyStatus).Methods("GET")
    r.HandleFunc("/api/errors", handlers.GetErrorResponses).Methods("GET")

    err := utils.OpenDatabase()
    if err != nil {
        log.Println("Error connecting to the database %v", err)
    }
    defer utils.CloseDatabase()
    
    log.Fatal(http.ListenAndServe(":8080", r))
}
