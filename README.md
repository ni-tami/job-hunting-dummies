# job-hunting-dummies
A job search app for dummy HR and dummy job hunters

## Entities
[DB Diagram file](./service/docs/diagram)

- User
> Id  
> Username  
> Name  
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Company
> Id  
> User  
> CompanyName  
> Description  
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Applicant
> Id  
> User  
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Application
> Id  
> Applicant  
> Job  
> Status
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Job
> Id  
> Company  
> Title  
> Description  
> Requirements  
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Requirements
> Requirement

- Status
> APPLIED > REVIEWED > PENDING_INTERVIEW > INTERVIEW_SCHEDULED > PENDING_RESULT > ACCEPTED / REJECTED
