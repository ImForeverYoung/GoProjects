
package model

type ServerResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}


type StudentDetailResponse struct {
	ID        int    `json:"id"`          
	FullName  string `json:"full_name"`   
	Gender    string `json:"gender"`      
	BirthDate string `json:"birth_date"`  
	GroupName string `json:"group_name"`  
}


type ScheduleResponse struct {
	GroupName  string `json:"group_name"`
	Subject string `json:"subject"`
	TimeSlot string `json:"time_slot"`
}