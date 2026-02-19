# job-hunting-dummies
A job search app for dummy HR and dummy job hunters

## Entities
[DB Diagram file](./service/docs/diagram)

- User
> ID  
> Username  
> Name  
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Company
> ID  
> User  
> CompanyName  
> Description  
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Applicant
> ID  
> User  
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Application
> ID  
> Applicant  
> Job  
> Status
> CreatedAt  
> UpdatedAt  
> DeletedAt  

- Job
> ID  
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
