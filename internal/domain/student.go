package domain

type Student struct {
	ID      int      `json:"id"`
	Email   string   `json:"email" binding:"required,email,max=64"`
	Name    string   `json:"name" binding:"required,min=3,max=64"`
	GPA     float64  `json:"gpa" binding:"required"`
	Courses []string `json:"courses"`
}

type UpdateStudentInput struct {
	Email   string   `json:"email" binding:"email,max=64"`
	Name    string   `json:"name" binding:"min=3,max=64"`
	GPA     float64  `json:"gpa"`
	Courses []string `json:"courses"`
}

type CreateStudentInput struct {
	Email   string   `json:"email" binding:"required,email,max=64"`
	Name    string   `json:"name" binding:"required,min=3,max=64"`
	GPA     float64  `json:"gpa" binding:"required"`
	Courses []string `json:"courses"`
}
